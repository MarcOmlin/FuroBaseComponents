import {LitElement, html} from 'lit-element';
import {FBP} from '@furo/fbp';
import './api-fetch.js';
import {CollectionControls} from "./lib/CollectionControls.js"

/**
 * `data-collection`
 *
 * an interface component to get collection requests
 * give this component a mime type of a resource and the hateoas. other additionally options like order-by filter could
 * also be added. based on those information the component generate a request and update the result to the collection
 *
 * How a hateoasIn likes:
 * {
 *  {rel: "self", method: "GET", type: "application/vnd.com.acme.task", href: "/api/v1/task"}
 *  {rel: "subresource", method: "GET", type: "application/vnd.com.acme.tag", href: "/api/v1/tag"}
 * }
 *
 * How to use:
 * ```
 * <data-collection ƒ-hts-in="--x"
 *                      service=""
 *                      ƒ-prev-page=""
 *                      ƒ-next-page=""
 *                      ƒ-last-page=""
 *                      ƒ-first-page=""
 *                      ƒ-list="--triggerGetCollection"
 *                      @-hts-out="--x"
 *                    >
 * </data-collection>
 * ```
 *
 * @customElement
 * @polymer
 * @appliesMixin FBP
 * @demo demo/collection.html
 */
class collectionAgent extends FBP(LitElement) {

  constructor() {
    super();
    this._servicedefinitions = window.Env.services;
    this._ApiEnvironment = window.Env.api;

    // HTS aus response anwenden
    this._FBPAddWireHook("--responseParsed", (r) => {
      if (Array.isArray(r.links)) {
        this.htsIn(r.links);
        /**
         * @event response-hts-updated
         * Fired when hateoas is updated from response
         * detail payload: {Array|HATEOAS}
         */
        let customEvent = new Event('response-hts-updated', {composed: true, bubbles: false});
        customEvent.detail = r.links;
        this.dispatchEvent(customEvent);
      }
    });
  }

  static get properties() {
    return {
      /**
       * Name des Services
       */
      service: {type: String, attribute: true},

      pageSize: {type: Number, attribute: "page-size"},
      fields: {type: String, attribute: true},
      orderBy: {type: String, attribute: "order-by"},
      filter: {type: Array, attribute: true},
      view: {type: String, attribute: true},
    };
  }

  firstUpdated() {
    super.firstUpdated();
    /**
     * @event collection-controls
     * Fired when
     * detail payload:
     */
    let customEvent = new Event('collection-controls', {composed: true, bubbles: true});

    customEvent.detail = new CollectionControls(this, this.service, this._servicedefinitions);
    setTimeout(() => {
      this.dispatchEvent(customEvent);
    }, 0);
  }


  /**
   * https://cloud.google.com/apis/design/design_patterns
   */

  /**
   * Partielle Repräsentation
   * https://cloud.google.com/apis/design/design_patterns#partial_response
   *
   * etwas seltsam, aber google sieht hier $fields vor
   *
   */
  setFields(fields) {
    this.fields = fields;
  }

  /**
   * Sortierreihenfolge
   * https://cloud.google.com/apis/design/design_patterns#sorting_order
   *
   * To avoid sql injection errors we do not send sql like syntax!
   *
   * order-by="foo,-bar"  means foo asc and bar desc
   */
  setOrderBy(order) {
    this.orderBy = order;
  }


  // Filtern  [["user","eq","12345"], ["abgeschlossen","eq", true]]
  setFilter(filterstring) {
    if (Array.isArray(filterstring)) {
      this.filter = filterstring
    }
  }

  // Gewünschte Seite. Tipp: Folge dem HATEOAS

  // Seitengrösse  page_size

  // Meta für die Anzahl der Elemente der Resource


  /**
   * contextbezogene Darstellung
   *
   * https://cloud.google.com/apis/design/design_patterns#resource_view
   *
   * view=smallcards
   *
   */


  /**
   * Setze den Service
   * @param service
   */
  set service(service) {
    if (!this._servicedefinitions[service]) {
      console.error("service " + service + " does not exist", this, "Available Services:", this._servicedefinitions);
      return;
    }
    this._service = this._servicedefinitions[service];

    if (this._service.general.lifecycle.deprecated) {
      console.warn("You are using a deprecated service (" + service + ") " + this._service.general.lifecycle.info);
    }
    // set pagination defaults
  }

  bindRequestObject(entityTree) {
    this._entityTree = entityTree;
  }


  _makeRequest(link, body) {
    let data;
    if (body) {
      data = JSON.stringify(body)
    }
    // Daten
    let headers = new Headers(this._ApiEnvironment.headers);
    headers.append('Content-Type', 'application/' + link.type + '+json');
    headers.append('Content-Type', 'application/json');

    let params = {};
    let r = link.href.split("?")
    let req = r[0];
    // add existing params
    if (r[1]) {
      r[1].split("&").forEach((p) => {
        let s = p.split("=");
        params[s[0]] = s[1];
      });
    }

    // Fields
    if (this.fields) {
      params.$fields = this.fields.split(' ').join('');
    }

    // Sort
    if (this.orderBy) {
      params.$order_by = this.orderBy.split(' ').join('');
    }

    // Filter
    if (this.filter) {
      params.$filter = JSON.stringify(this.filter);
    }

    // rebuild req
    let qp = [];
    for (let key in params) {
      qp.push(key + "=" + params[key]);
    }
    if (qp.length > 0) {
      req = req + "?" + qp.join("&");
    }

    return new Request(req, {
      method: link.method,
      headers: headers,
      body: data
    })
  }

  _checkServiceAndHateoasLinkError(rel, serviceName) {
    // check Service Get
    if (!this._service.services[serviceName]) {
      // todo fehler werfen ???
      console.warn("Restlet " + serviceName + " is not specified", this._service, this);
      return true;
    }

    // check Hateoas
    if (!this._hts[rel]) {
      console.warn("No HATEOAS for rel self", this._hts, this);
      return true;
    }
    return false;
  }

  _followRelService(rel, serviceName) {
    if (this._checkServiceAndHateoasLinkError(rel, serviceName)) {
      return;
    }
    this._FBPTriggerWire("--triggerLoad", this._makeRequest(this._hts[rel]));
  }

  /**
   * loads the entity if hts is available
   */
  list() {
    this._followRelService("list", "List");
  }


  /**
   * loads the entity if hts is available
   */
  first() {
    this._followRelService("first", "List");
  }


  /**
   * loads the entity if hts is available
   */
  prev() {
    this._followRelService("prev", "List");
  }


  /**
   * loads the entity if hts is available
   */
  next() {
    this._followRelService("next", "List");
  }

  /**
   * loads the entity if hts is available
   */
  last() {
    this._followRelService("last", "List");
  }


  htsIn(hts) {
    if (hts && hts[0] && hts[0].rel) {
      this._hts = {};
      hts.forEach((link) => {
        this._hts[link.rel] = link
      });
      /**
       * @event hts-updated
       * Fired when hateoas is updated
       * detail payload: Hateoas links
       */
      let customEvent = new Event('hts-updated', {composed: true, bubbles: false});
      customEvent.detail = hts;
      this.dispatchEvent(customEvent);
    }
  }

  render() {
    // language=HTML
    return html`
      <!-- Add a style block here -->
      <style>
        :host {
          display: none;
        }
      </style>
      <api-fetch
              ƒ-invoke-request="--triggerLoad"
              ƒ-abort-request="--abort-demanded"
              @-response="--responseParsed">
      </api-fetch>
    `;
  }

}

customElements.define('collection-agent', collectionAgent);
