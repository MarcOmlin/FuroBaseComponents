import {LitElement, html, css} from 'lit-element';
import {Theme} from "@furo/framework/theme"
import {FBP} from "@furo/fbp";
import "./furo-markdown"
import "./furo-doc/furo-doc-properties"
import "./furo-doc/furo-doc-methods"
import "./furo-doc/furo-doc-events"

/**
 * `furo-doc-element`
 * Renders a analysis file
 *
 * @summary todo shortdescription
 * @customElement
 * @demo demo/furo-doc-element.html
 * @appliesMixin FBP
 */
class FuroDocElement extends FBP(LitElement) {

  constructor() {
    super();
    this.element = {};


  }

  print(analysisElement) {
    this.element = analysisElement;
    this._FBPTriggerWire("--data", this.element);
    this.requestUpdate();
  }


  /**
   * flow is ready lifecycle method
   */
  __fbpReady() {
    super.__fbpReady();
    //this._FBPTraceWires()

  }

  /**
   * Themable Styles
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return Theme.getThemeForComponent(this.name) || css`
        :host {
            display: block;
            font-weight: 400;
            font-size: 14px;
        }

        :host([hidden]) {
            display: none;
        }
    `
  }


  /**
   * @private
   * @returns {TemplateResult}
   */
  render() {
    // language=HTML

    return html`
      <h1>${this.element.name}</h1>
      
      <furo-markdown ƒ-parse-markdown="--data(*.description)"></furo-markdown>
      <furo-doc-properties ƒ-data="--data(*.properties)"></furo-doc-properties>
      <furo-doc-methods ƒ-data="--data(*.methods)"></furo-doc-methods>
      <furo-doc-events ƒ-data="--data(*.events)"></furo-doc-events>
    `;
  }
}

window.customElements.define('furo-doc-element', FuroDocElement);