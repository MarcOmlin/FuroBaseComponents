import {LitElement, html, css} from 'lit-element';
import {Theme} from "@furo/framework/theme.js"
import {FBP} from "@furo/fbp";

/**
 * `furo-navigation-pad` listens to different keyboard navigation events like the arrow keys. It will attach the listeners
 *  to the parent node and cancel the default and stop the propagation of the events.
 *
 *  The events are available as standalone events or as combined event in the `navigated` event.
 *
 *  ```html
 *  <!-- forward all navigation events except the Escape  -->
 *  <furo-navigation-pad ignored-keys="Escape" @-navigated="--navpad"></furo-navigation-pad>
 *
 *  ```
 *
 *
 *
 *
 * @summary keyboard navigation helper
 * @customElement
 * @appliesMixin FBP
 */
class FuroNavigationPad extends FBP(LitElement) {


  /**
   * @private
   * @return {Object}
   */
  static get properties() {
    return {
      /**
       * Enter the keys you want to ignore as comma seperated values.
       *
       * i.e. "Escape, ArrowLeft"
       */
      ignoredKeys: {type: String, attribute: "ignored-keys"}
    };
  }

  /**
   * flow is ready lifecycle method
   */
  _FBPReady() {
    super._FBPReady();
    // this._FBPTraceWires()
    this.parentNode.addEventListener("keydown", (event) => {
      let key = event.key || event.keyCode;

      // abort on modifier keys
      if (event.ctrlKey || event.shiftKey || event.altKey || event.metaKey) {
        return
      }

      // abort on ignoredKeys
      if(this.ignoredKeys && this.ignoredKeys.split(",").filter(k=>(k.trim() === key)).length > 0){
        return;
      }


      /**
       * @event navigated
       * Fired when one of the keys was pressed
       * detail payload: key
       */
      let navigatedEvent = new Event('navigated', {composed: true, bubbles: true});

      switch (key) {
        case "Enter":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event enter-pressed
           * Fired when Enter key was pressed
           * detail payload: keyboard event
           */
          let enterEvent = new Event('enter-pressed', {composed: true, bubbles: true});
          enterEvent.detail = event;
          this.dispatchEvent(enterEvent);
          break;

        case "ArrowDown":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event arrow-down-pressed
           * Fired when ArrowDown key was pressed
           * detail payload: keyboard event
           */
          let arrowDownEvent = new Event('arrow-down-pressed', {composed: true, bubbles: true});
          arrowDownEvent.detail = event;
          this.dispatchEvent(arrowDownEvent);
          break;

        case "ArrowUp":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event arrow-up-pressed
           * Fired when ArrowUp key was pressed
           * detail payload: keyboard event
           */
          let arrowUpEvent = new Event('arrow-up-pressed', {composed: true, bubbles: true});
          arrowUpEvent.detail = event;
          this.dispatchEvent(arrowUpEvent);
          break;

        case "ArrowLeft":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);

          /**
           * @event arrow-left-pressed
           * Fired when ArrowLeft key was pressed
           * detail payload: keyboard event
           */
          let arrowLeftEvent = new Event('arrow-left-pressed', {composed: true, bubbles: true});
          arrowLeftEvent.detail = event;
          this.dispatchEvent(arrowLeftEvent);
          break;

        case "ArrowRight":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event arrow-right-pressed
           * Fired when ArrowRight key was pressed
           * detail payload: keyboard event
           */
          let arrowRightEvent = new Event('arrow-right-pressed', {composed: true, bubbles: true});
          arrowRightEvent.detail = event;
          this.dispatchEvent(arrowRightEvent);
          break;

        case"Escape":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event escape-pressed
           * Fired when Escape key was pressed
           * detail payload: keyboard event
           */
          let escapeEvent = new Event('escape-pressed', {composed: true, bubbles: true});
          escapeEvent.detail = event;
          this.dispatchEvent(escapeEvent);
          break;


        case"PageUp":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event page-up-pressed
           * Fired when PageUp key was pressed
           * detail payload: keyboard event
           */
          let pageUpEvent = new Event('page-up-pressed', {composed: true, bubbles: true});
          pageUpEvent.detail = event;
          this.dispatchEvent(pageUpEvent);
          break;

        case"PageDown":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event page-down-pressed
           * Fired when PageDown key was pressed
           * detail payload: keyboard event
           */
          let pageDownEvent = new Event('page-down-pressed', {composed: true, bubbles: true});
          pageDownEvent.detail = event;
          this.dispatchEvent(pageDownEvent);
          break;

        case"Home":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event home-pressed
           * Fired when Home key was pressed
           * detail payload: keyboard event
           */
          let homeEvent = new Event('home-pressed', {composed: true, bubbles: true});
          homeEvent.detail = event;
          this.dispatchEvent(homeEvent);
          break;

        case"End":
          event.stopPropagation();

          navigatedEvent.detail = key;
          this.dispatchEvent(navigatedEvent);
          /**
           * @event end-pressed
           * Fired when End key was pressed
           * detail payload: keyboard event
           */
          let endEvent = new Event('end-pressed', {composed: true, bubbles: true});
          endEvent.detail = event;
          this.dispatchEvent(endEvent);
          break;
      }

    });
  }

  /**
   * Themable Styles
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return Theme.getThemeForComponent('FuroNavigationPad') || css`
      :host {
        display: none;
      }
    `
  }

}

window.customElements.define('furo-navigation-pad', FuroNavigationPad);