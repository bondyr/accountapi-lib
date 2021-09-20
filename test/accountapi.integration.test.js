const { form3apis } = require('../src/js/form3apis')
const organisation = form3apis.organisation('v1')
const accounts = organisation.accounts
// googleapis.blogger('v3).blogs.get

const chai = require('chai')
const expect = chai.expect
const sinon = require('sinon')

const myOrgUuid = 'cdb76e34-74a1-4901-8179-196e7b133d55'

const requestBodyMin = {
  organisation_id: myOrgUuid,
  attributes: {
    name: ['Samantha Holder'],
    country: 'GB'
  }
}

const fullReqBodyPostman1 = {
  organisation_id: myOrgUuid,
  attributes: {
    name: ['Samantha Holder'],
    country: 'GB',
    base_currency: 'GBP',
    bank_id: '400302',
    bank_id_code: 'GBDSC',
    account_number: '10000004',
    customer_id: '234',
    iban: 'GB28NWBK40030212764204',
    bic: 'NWBKGB42',
    account_classification: 'Personal'
  }
}
const fullReqBodyPostman2 = {
  organisation_id: myOrgUuid,
  attributes: {
    country: 'GB',
    base_currency: 'GBP',
    bank_id: '400302',
    bank_id_code: 'GBDSC',
    customer_id: '234',
    bic: 'NWBKGB42',
    name: ['Samantha Holder', 'Sammy Holdy'],
    alternative_names: ['Sam Holder'],
    account_classification: 'Personal',
    joint_account: false,
    account_matching_opt_out: false,
    secondary_identification: 'A1B2C3D4'
  }
}
const fullReqBodyPostman3 = {
  organisation_id: myOrgUuid,
  attributes: {
    country: 'GB',
    base_currency: 'GBP',
    name: ['James Bond'],
    bank_id: '400305',
    bank_id_code: 'GBDSC',
    bic: 'LHVBEE22',
    account_classification: 'Personal',
    private_identification: {
      birth_date: '1920-11-11',
      birth_country: 'GB',
      identification: 'MI6008',
      address: ['11 Up and Down Street'],
      country: 'GB',
      city: 'London'
    }
  },
  relationships: {
    master_account: {
      data: [
        {
          type: 'accounts',
          id: '8d1d94e7-2974-433b-9759-3de2a84eaf03'
        }
      ]
    }
  }
}

describe('account resource create test suite', () => {
  let account_id
  let version
  it('should successfully create account with minimal body', async () => {
    const res = await accounts.create(requestBodyMin)
    expect(res.status).to.equal(201)
  })
  it('should successfully create account with postman1', async () => {
    const res = await accounts.create(fullReqBodyPostman1)
    expect(res.status).to.equal(201)
  })
  it('should successfully create account with postman2', async () => {
    const res = await accounts.create(fullReqBodyPostman2)
    expect(res.status).to.equal(201)
  })
  it('should successfully create account with postman3', async () => {
    const res = await accounts.create(fullReqBodyPostman3)
    expect(res.status).to.equal(201)
    const resData = res.body.data
    account_id = resData.id
    version = resData.version
  })
  it('should successfully get account', async () => {
    const res = await accounts.fetch(account_id)
    expect(res.status).to.equal(200)
  })
  it('should successfully delete account', async () => {
    const res = await accounts.delete(account_id, version)
    expect(res.status).to.equal(204)
  })
})
