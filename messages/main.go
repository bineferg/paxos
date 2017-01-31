package main
 
import (
    "encoding/json"
    "log"
    "net/http"
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "sync"

    "github.com/gorilla/mux"
    "github.com/asaskevich/govalidator"
)
//Returned json
type MessageResponse struct {
    Msg        string   `json:"message,omitempty"`
}

//Received json
type MessageRequest struct {
    Msg     string  `json:"message"valid:"required,string"`
}
//Used for Error handling
type Error struct {
    ErrMsg   string  `json:"err_msg,omitempty"`
    ErrCode  int   `json:"err_code,omitempty"`
}
 
type Sha_256 struct {
    Digest  string `json:"digest,omitempty"`
}

//Map
var digests = make(map[Sha_256]MessageRequest)

//Mutex for global map
var mutex  = &sync.Mutex{}

//Json validator
func (req *MessageRequest) Validate(w http.ResponseWriter, r *http.Request) error{
    if v, err := govalidator.ValidateStruct(req); !v || err != nil {
        return errors.New("Bad Request")
    }
    return nil
}


//Consumes message and reutrns sha digest of that message
func CreateShaEndpoint(w http.ResponseWriter, req *http.Request) {
    if req.Body == nil {
        errorHandler(w, req, http.StatusBadRequest)
        return
    }
    var r MessageRequest
    _ = json.NewDecoder(req.Body).Decode(&r)
    err := r.Validate(w, req)
    if err != nil {
        errorHandler(w, req, http.StatusBadRequest)
        return
    }

    //hash digest
    h := sha256.New()
    h.Write([]byte(r.Msg))
    sha256_hash := hex.EncodeToString(h.Sum(nil))
    s := Sha_256{Digest: sha256_hash}
    mutex.Lock()
    digests[s] = r
    mutex.Unlock()
    json.NewEncoder(w).Encode(s)
}

//Takes in sha and returns associated message if exists
func CreateMessageResponseEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var sha Sha_256
    sha.Digest = params["digest"]
    m, ok := digests[sha]
    if !ok {
        errorHandler(w, req, http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(m)

}

//Can be used for any error handling type
//Only 404 and 400 implemented here.
func errorHandler(w http.ResponseWriter, req *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        e := Error{ErrMsg: "MessageResponse not found", ErrCode:http.StatusNotFound}
        json.NewEncoder(w).Encode(e)
    }
    if status == http.StatusBadRequest {
        e := Error{ErrMsg: "Cannot process your request", ErrCode:http.StatusBadRequest}
        json.NewEncoder(w).Encode(e)
    }
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/message", CreateShaEndpoint).Methods("POST")
    router.HandleFunc("/message/{digest}", CreateMessageResponseEndpoint).Methods("GET")
    log.Fatal(http.ListenAndServe(":8001", router))
}
