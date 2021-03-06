import { FuroTimeInput } from '@furo/input/src/furo-time-input.js';
import { UniversalFieldNodeBinder } from '@furo/data/src/lib/UniversalFieldNodeBinder.js';

/**
 * `furo-data-time-input` is a extension of furo-text-input which enables you to
 *  bind a entityObject field.
 *
 * The field can be of type string, google.type.TimeOfDay, furo.fat.String or any type with the signature
 * of the google.protobuf.StringValue (string must be in field `value`).
 *
 * Setting the attributes on the component itself, will override the metas from spec, fat labels, fat attributes.
 *
 * ### following labels of the furo.fat.String are supported by default:
 *
 * - 'error': state of input is error
 * - 'readonly': input is disabled
 * - 'required': input is required
 * - 'disabled': input is disabled
 * - 'pristine': data is not changed. it is pristine
 * - 'condensed': input has condensed display
 *
 * ### following attributes of the furo.fat.String are supported by default:
 *
 * - 'label': input label
 * - 'hint': input hint
 * - 'errortext': the error text of the input
 * - 'error-msg': the same as errortext
 * - 'pattern': the input regex pattern.
 * - 'min': The earliest date and time to accept. format according to ISO 8601. e.g. '2000-01-01T08:32Z'
 * - 'max': The latest date and time to accept. format according to ISO 8601. e.g. '2000-01-01T08:32Z'
 * - 'step': The stepping interval to use for this input, such as when clicking arrows on spinner controls or performing validation
 *
 * ### following constrains are mapped into the attributes of the furo.fat.String and presence in payload:
 *
 * - 'max': is mapped to 'max' attribute
 * - 'min': is mapped to 'min' attribute
 * - 'step': is mapped to 'step' attribute
 * - 'required': is mapped to 'required' attribute
 *
 * <sample-furo-data-time-input></sample-furo-data-time-input>
 *
 * Tags: input
 * @summary Bind a entityObject.field to a time input
 * @customElement
 * @demo demo-furo-data-time-input Data binding
 * @mixes FBP
 */
export class FuroDataTimeInput extends FuroTimeInput {
  /**
   * @event value-changed
   * Fired when value has changed from inside the input field.
   *
   * detail payload: {String} the text value
   *
   * Comes from underlying component furo-time-input. **bubbles**
   */

  /**
   * @event trailing-icon-clicked
   * Fired when the trailing icon was clicked
   *
   * detail payload: the value of the time input
   *
   * Comes from underlying component furo-time-input. **bubbles**
   */

  /**
   * @event leading-icon-clicked
   * Fired when the leading icon was clicked
   *
   * detail payload: the value of the time input
   *
   * Comes from underlying component furo-time-input. **bubbles**
   */

  constructor() {
    super();
    this.error = false;
    this.disabled = false;

    this._initBinder();
  }

  /**
   * inits the universalFieldNodeBinder.
   * Set the mapped attributes and labels.
   * @private
   */
  _initBinder() {
    this.binder = new UniversalFieldNodeBinder(this);

    // set the attribute mappings
    this.binder.attributeMappings = {
      label: 'label',
      hint: 'hint',
      'leading-icon': 'leadingIcon',
      'trailing-icon': 'trailingIcon',
      errortext: 'errortext',
      'error-msg': 'errortext',
      min: 'min',
      max: 'max',
      step: 'step',
    };

    // set the label mappings
    this.binder.labelMappings = {
      error: 'error',
      readonly: 'readonly',
      required: 'required',
      disabled: 'disabled',
      condensed: 'condensed',
    };

    this.binder.fatAttributesToConstraintsMappings = {
      max: 'value._constraints.max.is', // for the fieldnode constraint
      min: 'value._constraints.min.is', // for the fieldnode constraint
      step: 'value._constraints.step.is', // for the fieldnode constraint
      required: 'value._constraints.required.is', // for the fieldnode constraint
      'min-msg': 'value._constraints.min.message', // for the fieldnode constraint message
      'max-msg': 'value._constraints.max.message', // for the fieldnode constraint message
    };

    this.binder.constraintsTofatAttributesMappings = {
      min: 'min',
      max: 'max',
      step: 'step',
      required: 'required',
    };

    /**
     * check overrides from the used component, attributes set on the component itself overrides all
     */
    this.binder.checkLabelandAttributeOverrrides();

    // the extended furo-text-input component uses _value
    this.binder.targetValueField = '_value';

    // update the value on input changes
    this.addEventListener('value-changed', val => {
      let TimeOfDayValue = val.detail;
      if (this.binder.fieldNode) {
        if (
          this.binder.fieldNode._spec.type === 'google.type.TimeOfDay' ||
          (this.binder.fieldNode['@type'] &&
            this.binder.fieldNode['@type']._value.replace(/.*\//, '') === 'google.type.TimeOfDay')
        ) {
          TimeOfDayValue = this._convertStringToTimeOfDayObj(
            TimeOfDayValue,
            this.binder.fieldNode._value,
          );
        }
        // store tmpval to check against loop
        this.tmpval = TimeOfDayValue;

        if (JSON.stringify(this.binder.fieldValue) !== JSON.stringify(TimeOfDayValue)) {
          // update the value
          this.binder.fieldValue = TimeOfDayValue;
        }
      }

      // set flag empty on empty strings (for fat types)
      if (this.binder.fieldFormat === 'fat') {
        if (TimeOfDayValue) {
          this.binder.deleteLabel('empty');
        } else {
          this.binder.addLabel('empty');
        }
        // if something was entered the field is not empty
        this.binder.deleteLabel('pristine');
      }
    });
  }

  //  The format is "HH:mm", "HH:mm:ss" or "HH:mm:ss.SSS" where HH is 00-23, mm is 00-59, ss is 00-59, and SSS is 000-999.
  // convert date string ISO 8601 to object for google.type.Dates
  // eslint-disable-next-line class-methods-use-this
  _convertStringToTimeOfDayObj(str, obj) {
    const arr = str.split(/[:.]/, 4);
    // eslint-disable-next-line no-param-reassign
    obj.hours = Number(arr[0] || null);
    // eslint-disable-next-line no-param-reassign
    obj.minutes = Number(arr[1] || null);
    // eslint-disable-next-line no-param-reassign
    obj.seconds = Number(arr[2] || null);
    // eslint-disable-next-line no-param-reassign
    obj.nanos = Number(arr[3] || null);
    return obj;
  }

  /**
   * Sets the value for the field. This will update the fieldNode.
   * @param val
   */
  setValue(val) {
    this.binder.fieldValue = val;
  }

  /**
   * Bind a entity field to the time-input. You can use the entity even when no data was received.
   * When you use `@-object-ready` from a `furo-data-object` which emits a EntityNode, just bind the field with `--entity(*.fields.fieldname)`
   * @param {Object|FieldNode} fieldNode a Field object
   */
  bindData(fieldNode) {
    this.binder.bindField(fieldNode);
    if (this.binder.fieldNode) {
      /**
       * handle pristine
       *
       * Set to pristine label to the same _pristine from the fieldNode
       */
      if (this.binder.fieldNode._pristine) {
        this.binder.addLabel('pristine');
      } else {
        this.binder.deleteLabel('pristine');
      }
      // set pristine on new data
      this.binder.fieldNode.addEventListener('new-data-injected', () => {
        this.binder.addLabel('pristine');
      });
    }
  }

  // because we defined the property max, the setter from the parent needs to be updated
  set max(val) {
    super.max = val;
  }

  // because we defined the property min, the setter from the parent needs to be updated
  set min(val) {
    super.min = val;
  }

  // because we defined the property step, the setter from the parent needs to be updated
  set step(val) {
    super.step = val;
  }

  static get properties() {
    return {
      /**
       * set this to true to indicate errors
       */
      error: { type: Boolean, reflect: true },
      /**
       * Overrides the label text from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      label: {
        type: String,
        reflect: true,
      },
      /**
       * Overrides the pattern from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      pattern: {
        type: String,
        reflect: true,
      },
      /**
       * Overrides the step value from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      step: {
        type: String, // string, because "any" is also a valid step
        reflect: true,
      },
      /**
       * Overrides the required value from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      required: {
        type: Boolean,
        reflect: true,
      },
      /**
       * Overrides the hint text from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      hint: {
        type: String,
        reflect: true,
      },
      /**
       * Overrides the min value from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      min: {
        type: Number,
        reflect: true,
      },
      /**
       * Overrides the max value from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      max: {
        type: Number,
        reflect: true,
      },
      /**
       * Overrides the readonly value from the **specs**.
       *
       * Use with caution, normally the specs defines this value.
       */
      readonly: {
        type: Boolean,
        reflect: true,
      },
      /**
       * A Boolean attribute which, if present, means this field cannot be edited by the user.
       */
      disabled: {
        type: Boolean,
        reflect: true,
      },

      /**
       * Set this attribute to autofocus the input field.
       */
      autofocus: {
        type: Boolean,
      },
      /**
       * Icon on the left side
       */
      leadingIcon: {
        type: String,
        attribute: 'leading-icon',
        reflect: true,
      },
      /**
       * Icon on the right side
       */
      trailingIcon: {
        type: String,
        attribute: 'trailing-icon',
        reflect: true,
      },
      /**
       * html input validity
       */
      valid: {
        type: Boolean,
        reflect: true,
      },
      /**
       * The default style (md like) supports a condensed form. It is a little bit smaller then the default
       */
      condensed: {
        type: Boolean,
        reflect: true,
      },
      /**
       * Lets the placeholder always float
       */
      float: {
        type: Boolean,
        reflect: true,
      },
    };
  }
}

customElements.define('furo-data-time-input', FuroDataTimeInput);
