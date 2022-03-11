package v1

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"teswir-go/internal/usecase"
	"teswir-go/pkg/logger"
)

type productRoutes struct {
	pUseCase usecase.Product
	log      logger.Interface
}

func newProductRoutes(handler *mux.Router, p usecase.Product, l logger.Interface) {
	r := productRoutes{p, l}

	handler.HandleFunc("/api/v1/product/add", r.productAdd).Methods("POST")
}

func (p *productRoutes) productAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("maladis")
	id, err := uuid.FromString("c54f3cf0-ac2d-4aac-891c-bdb4f19929f7")
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
