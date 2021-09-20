const accounts = require('./accounts')
// const branches = require('./branches')

function organisation (ver) {
  switch (ver) {
    case 'v1':
      return {
        accounts: accounts.accountsV1
        // branches: branches.branchesV1
      }
      // case 'v2':
      //     return {
      //         accounts: accounts.accountsV2,
      //         branches: branches.branchesV2
      //     }
    default:
      throw new Error(`No version ${ver} found of organisation property`)
  }
}

module.exports.organisation = organisation
