/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package storage

/*****************************************************************************************************************/

import (
	pb "birpc/internal/gen/store/v1/storev1connect"
)

/*****************************************************************************************************************/

type server struct {
	pb.UnimplementedStorageServiceHandler
}

/*****************************************************************************************************************/

func NewStorageServer() *server {
	return &server{}
}

/*****************************************************************************************************************/
