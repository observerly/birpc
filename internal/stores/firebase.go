/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package stores

/*****************************************************************************************************************/

import (
	"firebase.google.com/go/v4/storage"
)

/*****************************************************************************************************************/

type FirebaseStore struct {
	Store
	Client *storage.Client
}

/*****************************************************************************************************************/

func NewFirebaseStorageClient(client *storage.Client) *FirebaseStore {
	return &FirebaseStore{
		Store:  Store{},
		Client: client,
	}
}

/*****************************************************************************************************************/
