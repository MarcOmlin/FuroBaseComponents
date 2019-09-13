// Code generated by @furo/specs. DO NOT EDIT.
// source: person.person.referencesearch.spec
import { LitElement, html, css } from 'lit-element';
import {Theme} from "@furo/framework/theme"
import {FBP} from "@furo/fbp";
import "@furo/data"
import "@furo/data-input"



/**
 * `person-person-reference-search`
 *  Complete reference searcher for person.Person
 *
 *  
 *
 * @summary search person.Person
 * @customElement
 * @appliesMixin FBP
 */
class PersonPersonReferenceSearch extends FBP(LitElement) {

  bindData(field){
    this._FBPTriggerWire("--field-injected", field);
  }
  /**
   * Themable Styles
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return Theme.getThemeForComponent('ReferenceBaseTheme') || css`
        :host {
            display: block;
        }

        :host([hidden]) {
            display: none;
        }
    `
  }


  /**
   * @private
   * @returns {TemplateResult}
   * @private
   */
  render() {
    // language=HTML
    return html`
      <furo-data-reference-search ƒ-bind-data="--field-injected"
                                  @-search="--term"
                                  ƒ-collection-in="--collection">
      </furo-data-reference-search>
      <furo-collection-agent
              service="ProjectMembersService"
              ƒ-hts-in="--field-injected(*.link.value)"
              ƒ-search="--term"
              @-response="--collection">
      </furo-collection-agent>
    `;
  }
}

window.customElements.define('person-person-reference-search', PersonPersonReferenceSearch);
