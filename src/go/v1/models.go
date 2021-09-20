// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.

package accounts

type AccountStruct struct {
	Data            AccountData       `json:"data,omitempty"`
	Links		   *AccountLinks      `json:"links,omitempty"`
}

type AccountLinks struct {
	Self           string             `json:"self,omitempty"`
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
	CreatedOn      *string            `json:"created_on,omitempty"`
	ModifiedOn     *string            `json:"modified_on,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

// 201 {
//   data: {
//     attributes: {
//       alternative_names: null,
//       bank_id: 'MYBANKID',
//       base_currency: 'PLN',
//       country: 'GB',
//       name: [Array]
//     },
//     created_on: '2021-09-30T19:29:10.552Z',
//     id: '96f7e899-8f70-45c6-b9e6-c5320b7de7d6',
//     modified_on: '2021-09-30T19:29:10.552Z',
//     organisation_id: 'cdb76e34-74a1-4901-8179-196e7b133d55',
//     type: 'accounts',
//     version: 0
//   },
//   links: {
//     self: '/v1/organisation/accounts/96f7e899-8f70-45c6-b9e6-c5320b7de7d6'
//   }
// }