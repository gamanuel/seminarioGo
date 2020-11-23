package cryptocurrency

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)
// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method string
	path string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}
	
	list = append(list, 
		&endpoint{
			method: "GET",
			path: "/cryptocurrency",
			function: getAll(s),
		},
		&endpoint{
			method: "GET",
			path: "/cryptocurrency/:id",
			function: getCryptocurrencyById(s),
		},
		&endpoint{
			method: "POST",
			path: "/cryptocurrency",
			function: addCryptocurrency(s),
		},
		&endpoint{
			method: "DELETE",
			path: "/cryptocurrency/:id",
			function: deleteCryptocurrency(s),
		},
		&endpoint{
			method: "PUT",
			path: "/cryptocurrency",
			function: updateCryptocurrency(s),
		},
)



	return list
}


func getAll(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		result, _ := s.FindAll()
		c.JSON(http.StatusOK, gin.H {
			"cryptocurrency": result,
		})
	}
}

func getCryptocurrencyById(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, _ := s.FindByID(ID)
		c.JSON(http.StatusOK, gin.H {
			"cryptocurrency": result,
		})
	}
	
}

func addCryptocurrency(s Service) gin.HandlerFunc {
	var cryptocurrency Cryptocurrency
	return func(c *gin.Context) {
		c.BindJSON(&cryptocurrency)
		result, _ := s.AddCryptocurrency(cryptocurrency)
		c.JSON(http.StatusOK, gin.H{
			"cryptocurrency": result,
		})
	}
}

func deleteCryptocurrency(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, _ := strconv.Atoi(c.Param("id"))
		result, _ := s.RemoveByID(ID)
		c.JSON(http.StatusOK, gin.H {
			"cryptocurrency": result,
		})
	}
}



func (s httpService) Register(r *gin.Engine){
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

func updateCryptocurrency(s Service) gin.HandlerFunc {
	var cryptocurrency Cryptocurrency
	return func(c *gin.Context) {
		c.BindJSON(&cryptocurrency)
		result, _ := s.updateCryptocurrency(cryptocurrency)
		c.JSON(http.StatusOK, gin.H{
			"cryptocurrency": result,
		})
	}
}
