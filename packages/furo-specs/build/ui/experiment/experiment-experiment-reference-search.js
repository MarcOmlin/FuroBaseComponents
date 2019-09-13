// Code generated by @furo/specs. DO NOT EDIT.
// source: experiment.experiment.referencesearch.spec
import { LitElement, html, css } from 'lit-element';
import {Theme} from "@furo/framework/theme"
import {FBP} from "@furo/fbp";
import "@furo/data"
import "@furo/data-input"



/**
 * `experiment-experiment-reference-search`
 *  Complete reference searcher for experiment.Experiment
 *
 *  
 *
 * @summary search experiment.Experiment
 * @customElement
 * @appliesMixin FBP
 */
class ExperimentExperimentReferenceSearch extends FBP(LitElement) {

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
              service="ExperimentService"
              ƒ-hts-in="--field-injected(*.link.value)"
              ƒ-search="--term"
              @-response="--collection">
      </furo-collection-agent>
    `;
  }
}

window.customElements.define('experiment-experiment-reference-search', ExperimentExperimentReferenceSearch);
