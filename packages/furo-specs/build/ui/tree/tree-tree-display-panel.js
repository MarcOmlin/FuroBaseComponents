// Code generated by @furo/specs. DO NOT EDIT.
// source: ./specs/tree/tree.service.spec
import {html, css} from 'lit-element';
import {FBP} from "@furo/fbp";
import {Theme} from "@furo/framework/theme"
import {i18n} from "@furo/framework/i18n"
import {BasePanel} from "@furo/route/lib/BasePanel";

import '@furo/data';
import '@furo/layout';


import "./tree-tree-display";

/**
 * The Get method takes zero or more parameters, and returns a TreeEntity which contains a Tree
 *
 * @customElement
 * @appliesMixin FBP
 */
export class TreeTreeDisplayPanel extends FBP(BasePanel) {

    static get styles() {
        // language=CSS
        return Theme.getThemeForComponent('PanelBaseTheme') || css`
                :host {
                    display: block;
                    height: 100%;
                    background-color: var(--surface-light);
                    color: var(--on-surface);
                    padding-top: var(--spacing);
                }

                :host([hidden]) {
                    display: none;
                }

                furo-card {
                    margin: 0 var(--spacing);
                    margin-bottom: var(--spacing);
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
          <furo-vertical-flex>
            <furo-card>

              <tree-tree-display flex ƒ-bind-data="--entity(*.data)"></tree-tree-display>
            </furo-card>


          </furo-vertical-flex>


          <furo-entity-agent service="TreeService"
                             @-response="--response"
                             ƒ-hts-in="--navNode(*.value.link)"
                             ƒ-bind-request-data="--entity(*.data)"
                             load-on-hts-in></furo-entity-agent>

          <furo-data-object type="tree.TreeEntity"
                            @-object-ready="--entity"
                            ƒ-reset="--resetReq"
                            ƒ-inject-raw="--response"></furo-data-object>

        `;
    }


}

window.customElements.define('tree-tree-display-panel', TreeTreeDisplayPanel);