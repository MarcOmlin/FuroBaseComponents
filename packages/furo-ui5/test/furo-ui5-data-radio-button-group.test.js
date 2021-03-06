import { fixture, html } from '@open-wc/testing';
import 'axe-core/axe.min.js';
import { axeReport } from 'pwa-helpers/axe-report.js';
// eslint-disable-next-line import/no-extraneous-dependencies
import '@furo/data/src/furo-data-object.js';
import '@furo/fbp/src/testhelper/test-bind.js'; // for testing with wires and hooks
// eslint-disable-next-line import/no-extraneous-dependencies
import '@furo/testhelper/initEnv.js';

import '../src/furo-catalog.js';

describe('furo-ui5-data-radio-button-group', () => {
  let host;
  let dropdown;
  let radioGroup;
  let dao;

  function keydown(TargetElement, key) {
    const customEvent = new Event('keydown', { composed: true, bubbles: true });
    customEvent.code = key; // Deprecated, prefer .key instead.
    customEvent.key = key;
    TargetElement.dispatchEvent(customEvent);
  }

  const testData = {
    entities: [
      {
        data: {
          display_name: 'John Doe, +41783332244',
          first_name: 'John',
          id: '1',
          name: 'Doe',
          phone_nr: '+41783332244',
          skills: [],
        },
        links: [
          {
            href: '/mockdata/persons/1/get.json',
            method: 'GET',
            rel: 'self',
            type: 'person.Person',
            service: 'PersonService',
          },
        ],
      },
      {
        data: {
          display_name: 'Tari Sakota, +41791532244',
          first_name: 'Tari',
          id: '2',
          name: 'Sakota',
          phone_nr: '+41791532244',
          skills: [],
        },
        links: [
          {
            href: '/mockdata/persons/2/get.json',
            method: 'GET',
            rel: 'self',
            type: 'person.Person',
            service: 'PersonService',
          },
        ],
      },
      {
        data: {
          display_name: 'Yoko Tasimoto, +41781442244',
          first_name: 'Yoko',
          id: '3',
          name: 'Tasimoto',
          phone_nr: '+41781442244',
          skills: [],
        },
        links: [
          {
            href: '/mockdata/persons/3/get.json',
            method: 'GET',
            rel: 'self',
            type: 'person.Person',
            service: 'PersonService',
          },
        ],
      },
      {
        data: {
          display_name: 'Lola Tasimoto, +41781442244',
          first_name: 'Lola',
          id: '4',
          name: 'Tasimoto',
          phone_nr: '+41781442244',
          skills: [],
        },
        links: [
          {
            href: '/mockdata/persons/4/get.json',
            method: 'GET',
            rel: 'self',
            type: 'person.Person',
            service: 'PersonService',
          },
        ],
      },
    ],
    links: [
      {
        href: '/mockdata/persons/list.json',
        method: 'GET',
        rel: 'list',
        type: 'person.PersonCollection',
        service: 'PersonService',
      },
    ],
  };
  const testDataArray = [
    {
      id: 1,
      display_name: 'Item 1',
    },
    {
      id: 2,
      display_name: 'Item 2',
    },
    {
      id: 3,
      display_name: 'Item 3',
    },
  ];

  beforeEach(async () => {
    const testbind = await fixture(html`
      <test-bind>
        <template>
          <furo-ui5-data-radio-button-group
            ƒ-bind-data="--typeData(*.owner)"
            sub-field="data"
          ></furo-ui5-data-radio-button-group>
          <furo-ui5-data-text-input ƒ-bind-data="--typeData(*.owner.id)"></furo-ui5-data-text-input>
          <furo-data-object type="task.Task" @-object-ready="--typeData"></furo-data-object>
        </template>
      </test-bind>
    `);
    await testbind.updateComplete;
    host = testbind._host;
    [, dropdown, radioGroup, dao] = testbind.parentNode.children;
    await host.updateComplete;
    await dropdown.updateComplete;
    await radioGroup.updateComplete;
    await dao.updateComplete;
  });

  it('should be a furo-ui5-data-radio-button-group element', done => {
    // keep this test on top, so you can recognize a wrong assignment
    assert.equal(dropdown.nodeName.toLowerCase(), 'furo-ui5-data-radio-button-group');
    done();
  });

  it('should have the correct items', done => {
    assert.equal(dropdown._dropdownList.length, 3);
    done();
  });
});
