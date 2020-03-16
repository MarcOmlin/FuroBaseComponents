import {LitElement, html, css} from 'lit-element';
import {Theme} from "@furo/framework/theme"
import {FBP} from "@furo/fbp";
import "@furo/doc-helper";
import "@furo/app/furo-card.js";
import "@furo/input/furo-button.js";
import "../furo-catalog";

/**
 * `demo-furo-file-dialog`
 *
 * @customElement
 * @appliesMixin FBP
 */
class DemoFuroFileDialog extends FBP(LitElement) {

    /**
     * Themable Styles
     * @private
     * @return {CSSResult}
     */
    static get styles() {
        // language=CSS
        return Theme.getThemeForComponent('DemoFuroFileInput') || css`
            :host {
                display: block;
                height: 100%;
                padding-right: var(--spacing);
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
            <h2>Demo furo-file-input</h2>
            <p>let the user choose one or more files from their device storage. </p>
            <furo-demo-snippet style="height: 600px">
                <template>
                    <furo-vertical-scroller>
                    <div style="display: grid; padding: var(--spacing); background-color: var(--surface-dark); grid-row-gap: var(--spacing); grid-column-gap: var(--spacing); grid-template-columns: repeat(3, 1fr);">
                        <furo-card header-text="Choose your media">
                           
                            <li>Accept: .jpg</li>
                            <li>Multiple: false</li>
                            <li>Required: false</li>
                            <div ƒ-.inner-text="--selected(*.detail.files.0.name)"></div>
                            <furo-button slot="action" label="Upload image" outline primary @-click="--openDlg"></furo-button>
                            <furo-file-dialog accept=".jpg" ƒ-open="--openDlg" @-input-changed="--selected(*)"></furo-file-dialog>

                        </furo-card>
                        <furo-card header-text="Choose your file">
                            <li>Accept: .md</li>
                            <li>Multiple: false</li>
                            <li>Required: false</li>
                            <li>Disabled</li>
                            <furo-button slot="action" label="Upload image" outline primary disabled @-click="--openDlg2"></furo-button>
                            <furo-file-dialog accept=".md" ƒ-open="--openDlg2"></furo-file-dialog>

                        </furo-card>

                    </div>
                    </furo-vertical-scroller>
                </template>
            </furo-demo-snippet>
        `;
    }
}

window.customElements.define('demo-furo-file-dialog', DemoFuroFileDialog);