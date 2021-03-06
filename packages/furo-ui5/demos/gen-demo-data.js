import { LitElement } from 'lit-element';

/**
 * `gen-demo-data`
 * todo Describe your element
 *
 * @summary todo shortdescription
 * @customElement
 * @demo demo-gen-demo-data
 * @appliesMixin FBP
 */
class GenDemoData extends LitElement {
  generate() {
    const data = {
      data: {
        id: '1',
        scalar_string: 'this is a scalar string',
        wrapper_string: {
          value: 'this is a google wrapper string',
        },
        fat_string: {
          value: 'this is a demo fat string',
          labels: ['required', 'condensed', 'error'],
          attributes: {
            'value-state': 'Error',
            'error-msg': 'XXXXX something went wrong',
            errortext: 'something went wrong',
            hint: 'with error and icon',
            'leading-icon': 'filter',
            label: 'override and required',
            placeholder: 'Placeholder override and required',
            max: 2,
          },
        },

        fat_string_list: {
          value: 'list',
          labels: [],
          attributes: {
            'value-state': 'Information',
            'information-msg': 'Max isset to 6',
            max: 6,
            'max-msg': 'max 6 symbols',
          },
        },
        scalar_int32: 14,
        wrapper_int32: {
          value: 14,
        },
        fat_int32: {
          value: 14,
          'value-state': 'Information',
        },
      },
      links: [],
      meta: {
        fields: {
          'data.wrapper_string': {
            meta: {
              label: 'wrapper string label setted via response meta',
              readonly: false,
              default: 'new',
            },
          },
          'data.scalar_string': {
            meta: {
              label: 'scalar_string string label setted via response meta',
              readonly: false,
              default: 'new',
            },
          },
          'data.fat_string': {
            meta: {
              label: 'fat string label setted via response meta',
              readonly: false,
              default: 'new',
            },
          },
        },
      },
    };
    /**
     * @event data
     * Fired when data is generated
     * detail payload: universaltest.Entity
     */
    const customEvent = new Event('data', { composed: true, bubbles: true });
    customEvent.detail = data;
    this.dispatchEvent(customEvent);
  }

  /**
   * @event value-state-list
   * Is fired if the valueStateList is generated
   * Available options are:
   * None
   * Error
   * Warning
   * Success
   * Information
   */
  generateValueStateList() {
    const list = [
      {
        id: '1',
        display_name: 'None',
      },
      {
        id: '2',
        display_name: 'Error',
      },
      {
        id: '3',
        display_name: 'Warning',
      },
      {
        id: '4',
        display_name: 'Success',
      },
      {
        id: '5',
        display_name: 'Information',
      },
    ];

    const customEvent = new Event('value-state-list', { composed: true, bubbles: true });
    customEvent.detail = list;
    this.dispatchEvent(customEvent);
  }
}

window.customElements.define('gen-demo-data', GenDemoData);
