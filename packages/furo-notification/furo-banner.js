import {LitElement,html,css} from 'lit-element';

/**
 * `furo-banner`
 * Lit element
 *
 *  furo-banner should be used together witch furo-banner-display. you can place those two components into different places.
 *  best place the furo-banner-display on the main site. then you only need one furo-banner-display. it can work with n furo-banner.
 *
 * @customElement
 * @demo demo-furo-banner-display banner demo
 */
class FuroBanner extends LitElement {


  constructor(){
    super();
    this.dismissButtonText = "dismiss";
    //this.confirmButtonText = "confirm";
  }


  /**
   * @private
   * @returns {CSSResult}
   */
  static get styles() {
    return css`
            :host {
            }
        `;
  }

  /**
   *@private
   */
  static get properties(){

    return {
      text: {
        type: String
      },
      dismissButtonText: {
        type: String,
        attribute: "dismiss-button-text"
      },
      confirmButtonText: {
        type: String,
        attribute: "confirm-button-text"
      },
      icon: {
        type: String
      },
      payload: {
        type: Object
      }
    };
  }

  setIcon(i) {
    this.icon = i;
  }

  setText(t) {
    this.text = t;
  }

  setConfirmButtonText(t) {
    this.confirmButtonText = t;
  }

  setDismissButtonText(t) {
    this.dismissButtonText = t;
  }

  /**
   * show banner
   * @param p {Object} payload
   */
  show(p) {

    this.payload = p;
    /**
     * @event open-furo-banner-requested
     * Fired when value open banner is requested
     * detail payload: {Object}  this
     */
    let customEvent = new Event("open-furo-banner-requested",{composed: true, bubbles: true});
    customEvent.detail = this;
    this.dispatchEvent(customEvent);
  }


  /**
   *  event `dismissed` will be sent with payload
   */
  dismiss() {

    /**
     * @event dismissed
     * Fired when dismiss button of banner is clicked
     * detail payload: {Object}  payload
     */
    let customEvent = new Event("dismissed",{composed: true, bubbles: true});
    customEvent.detail = this.payload;
    this.dispatchEvent(customEvent)
  }

  /**
   *  event `confirmed` will be sent with payload
   */
  confirm() {

    /**
     * @event confirmed
     * Fired when confirm button of banner is clicked
     * detail payload: {Object}  payload
     */
    let customEvent = new Event("confirmed",{composed: true, bubbles: true});
    customEvent.detail = this.payload;
    this.dispatchEvent(customEvent)
  }

  /**
   * parse grpc status object and set the label according to the message in status
   * @param s
   */
  parseGrpcStatus(s) {

    if (s.message) {
      this.setText(s.message);
      this.show(s);
    }
  }

  /**
   * @private
   * @returns {TemplateResult}
   */
  render(){
    return html`
        `;
  }

}

customElements.define('furo-banner', FuroBanner);
