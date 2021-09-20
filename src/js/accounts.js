'use strict'
'esversion: 8'

const superagent = require('superagent')
const uuidv4 = require('uuid').v4

const host = process.env.ACCOUNT_API_LIB_HOST || 'https://api.staging-form3.tech'
const endpoint = '/v1/organisation/accounts'
const baseUrl = `${host}${endpoint}`

async function createAccount (body) {
  const url = `${baseUrl}`
  try {
    const res = await superagent.post(url).send(body)
    console.error(res.status, res.body)
    return res
  } catch (error) {
    console.error(error.status, error.response.body)
    throw error
  }
}

function prepareCreateAccountBody (jsonOb) {
  const body = {
    data: {
      id: uuidv4(),
      type: 'accounts',
      ...jsonOb
    }
  }
  return body
}

const accountsV1 = {
  async fetch (account_id) {
    const url = `${baseUrl}/${account_id}`
    try {
      const res = await superagent.get(url)
      console.error(res.body)
      return res
    } catch (error) {
      console.error(error.status, error.response.body)
      throw error
    }
  },
  async create (jsonOb) {
    const url = `${baseUrl}`
    console.error(url)
    const body = prepareCreateAccountBody(jsonOb)
    try {
      const res = await superagent.post(url).send(body)
      console.error(res.status, res.body)
      return res
    } catch (error) {
      console.error(error.status, error.response.body)
      throw error
    }
  },
  async delete (account_id, version) {
    const url = `${baseUrl}/${account_id}?version=${version}`
    try {
      const res = await superagent.delete(url)
      console.error(res.body)
      return res
    } catch (error) {
      console.error(error.status, error.response.body)
      throw error
    }
  }
}

module.exports.accountsV1 = accountsV1
