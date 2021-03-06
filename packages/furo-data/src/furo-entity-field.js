import { LitElement } from 'lit-element';

/**
 * Use this component to interact with fields from an furo-data-object.
 *
 * You can set the field value or listen to changes of a field.
 *
 * @summary interact with single field of a data object
 * @customElement
 * @mixes FBP
 */
class FuroEntityField extends LitElement {
  /**
   * Set the value of the field.
   * @param v
   */
  setValue(v) {
    this.value = v;
  }

  set value(v) {
    this._value = v;
    this.field._value = v;
  }

  get value() {
    return this._value;
  }

  /**
   * Bind a entity field to the date-input. You can use the entity even when no data was received.
   * When you use `@-object-ready` from a `furo-data-object` which emits a EntityNode, just bind the field with `--entity(*.fields.fieldname)`
   * @param {Object|FieldNode} fieldNode a Field object
   */
  bindData(fieldNode) {
    if (fieldNode === undefined) {
      // eslint-disable-next-line no-console
      console.warn('Invalid binding ');
      // eslint-disable-next-line no-console
      console.log(this);
      return new Error('Invalid binding');
    }

    this.field = fieldNode;

    this.field.addEventListener('field-value-changed', e => {
      /**
       * @event value-changed
       * Fired when
       * detail payload:
       */
      const customEvent = new Event('value-changed', { composed: true, bubbles: true });
      customEvent.detail = e.detail.value;
      this.dispatchEvent(customEvent);
    });
    return this.field;
  }
}

customElements.define('furo-entity-field', FuroEntityField);
