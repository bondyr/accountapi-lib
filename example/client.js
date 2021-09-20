'use strict'

const { form3apis } = require('../src/js/form3apis')
const organisation = form3apis.organisation('v1')
const accounts = organisation.accounts

const organisation_id = 'cdb76e34-74a1-4901-8179-196e7b133d55'

const requestBody = {
  organisation_id,
  attributes: {
    name: ['Samantha Holder'],
    country: 'GB',
    base_currency: 'PLN',
    bank_id: 'MYBANKID'
  }
}

async function runApiLibFunctions () {
  console.error('============ Creating account ==========')
  const createRes = await accounts.create(requestBody)

  console.error('============ Fetching account ==========')
  const account_id = createRes.body.data.id
  console.log(account_id)
  await accounts.fetch(account_id)
  console.error('============ Deleting account ==========')
  const version = createRes.body.data.version
  await accounts.delete(account_id, version)
}

runApiLibFunctions()
