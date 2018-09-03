package photographer

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	storage "google.golang.org/api/storage/v1"
)

//TODO: Refactor Upload into an interface or something like that

func getGoogleCloud() (service *storage.Service, err error) {
	authConf := &jwt.Config{
		Email:      "starting-account-txi8an4ffyhb@yescort-212221.iam.gserviceaccount.com",
		PrivateKey: []byte("-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCv3VRFhxprnU3U\nkijfsclhotI9IPiLe1f0xEIhSLcwjTP6BCzXL3zhe8ifGg7+/uI3kQd5pJM9kce0\nHbQBC/e7EFlq58mk5GH+IukivtOeMgukEBEQjbEmVsNVjdvlg70IkohPkMfZ1WOB\n1ehVDIMgCOCeNInOUseez679hXGJTsS2xEdkNCe7CwC9OCLbkb6qp0F++cRMmM7z\nKf3ISHr+VYXFCXtut/agUfJGGb7qdNwwp5C65Bl6Iy5oVwpngcwN/oYXzyB6RdeW\n1AuVKbxeHEDdw6czYSipB9sAbu8Hf+EUFmXFTmu5j5+sYUUpkFWbu1/W9uc0xTM3\nE/xfAo3PAgMBAAECggEAGv+CaW4VO0IZ6bFKfxTfEpMZbXL7VQZpeVULxepQ3S5A\nNXignaRXqpIDOZGZjcH5XohwOldulzwjCV6p/KMms9wNriDJTvwUoxwe6EudVyLd\nsKEazYzDDlyxfhfaE0EHKyK7tGNImwNsVElPcocKNkfGm+L66ObCaHn/pkHvLSHd\nyHFYUF/Pu4Fp5sXXPbU72w1a3QDxybFn1N98PjkcqrJ0WhehdyeFr4iYOT1yFfBh\nkhwUW6OMJ9/dIeRXTrBp19MNnQprVupuMfYVLPvXeNyIq2DuTSzZbWaPiDrp8/eW\nBP+IG9HDi2is74kqo41pe9W0veNePzo3+Vs+sQQL8QKBgQDy6yUV3CQ5fEUyXfgP\noM4XPfb68tBxpO6pbbr/sOlBIyZTg5/p/krQSmqqiZrPe/rEwNrW1r0PuHkRthqb\nrxjP+oVOy2NnrzKA878baqiBTUZ4FGzzkkCKd3mxeUfpqUR8SRI7l2PZnovfnl3+\nGRtcBDuAcVmWA6x3a16P/ArflwKBgQC5VcimM/bJ2eFgwE2jB46p4quymjY6iuG0\nYvknn+ZQr12Hr2kfEgBGjBf820RaZEL9mQd3XDqVvjQZV/8ujD67l8n9+BILdHEg\nBcZpfC6gFtCA9UQ8pZoTdlU371zoiswsxjF0KfM89n7rbsbXfvzf6WbE+FoQKkp/\nmqvHUdMKiQKBgQDW6k+PPYsXCOk3nIehXAW57mqar0TXprI4WAq/uUdFB9IaxTMw\nFHkgOxaP9tpugyPj+zQiKy5twCntD44O+yUarGxQwAUccJZgHEaR/RVEAJg+GOQj\nwtszJm0WJVUIV2duBkAJVPFUUb8ygqxp6NCZ4my6QI6f+HQX23h4Lud31QKBgDAi\nbHtmEZl5IgTbzQ4jf1T2DvvS+3SW72nHLsiwycN42ETONPIh3bp01iIQ3J7RXWi5\n3uuMI5cSkDto1FjFieF02bxBOoaJFnxno5Fv9DgV1M85ZCCQLgUR/a4KV8n5im+y\nhdOStZ7Gsk35NmFvlMOevdc4XUPHiBO+GxEoWjRpAoGBALxQrpeVU2u1RR8T7LQd\n1ybkD0+N8MwhJNW9pETWz0QVy++qP0yESSriaX96vCOZYKzEvaUbycMb5uoZ7NPN\njja/12mVMNaRsv8VTxdAv8LdriXhXPHdkdbWnp1Gy3m8KUbc+mhQFyxuPmoIC5R1\nHj39v3YwqogtU3Ju601wWrBf\n-----END PRIVATE KEY-----\n"),
		Scopes:     []string{storage.DevstorageReadWriteScope},
		TokenURL:   "https://accounts.google.com/o/oauth2/token",
	}

	client := authConf.Client(oauth2.NoContext)
	service, err = storage.New(client)
	if err != nil {
		return service, errors.New("Problem authenticating to GCS")
	}
	return service, nil
}

func uploadProfilePicture(r *http.Request) (url string, err error) {
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		fmt.Printf("parseForm: %v\n", err)
		return "", nil
	}

	file, fh, err := r.FormFile("avatar")
	if err == http.ErrMissingFile {
		fmt.Printf("Missing File %v", err)
		return "", nil
	}
	if err != nil {
		fmt.Printf("Error: %v", err)
		return "", err
	}
	fmt.Println("After Error checking")

	// random filename, retaining existing extension.
	name := fmt.Sprintf("%s_%s%s", uuid.Must(uuid.NewV4()).String(), fh.Filename, path.Ext(fh.Filename))

	service, err := getGoogleCloud()
	if err != nil {
		log.Println(err)
	}

	object := &storage.Object{
		Name:         name,
		CacheControl: "public, max-age=31536000",
		Acl:          []*storage.ObjectAccessControl{{Entity: "allUsers", Role: "READER"}},
	}
	w, err := service.Objects.Insert("yescort", object).Media(file).Do()
	if err != nil {
		return "", err
	}
	fmt.Printf("%v", w)
	//w.Acl = []*storage.ObjectAccessControl{{Entity: "allUsers", Role: "READER"}}

	const publicURL = "https://storage.googleapis.com/%s/%s"
	return fmt.Sprintf(publicURL, "yescort", name), nil
}
