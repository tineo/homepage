package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"net/smtp"


)

type Tineo struct {
	FirstName string
	LastName string
	Email string
}

type Person struct {
	Name string
	Email string
}

type Education struct {
	Logo string
	College string
	Career string
	Lapse string
}

type Work struct {
	Logo string
	Company string
	BussinessName string
	Department string
	Lapse string
}

type Skills struct {
	Logo string
	Color string
	SkillName string
	Percent int8
}

type Services struct {
	Logo string
	Color string
	ServiceName string
	Description string
}

type Portfolio struct {
	Title string
	ImageCaption string
	Categories string
	LongTitle string
	BigImageCaption string
	Description string
	Url string
}

type References struct {
	Image string
	Description string
}

type Contact struct {
	Message string
	Email string
	Cellphone string
	Skype string
	Address string
}

type Social struct {
	Icon string
	Url string
	Caption string
}

type Personal struct {
	FirstName string
	LastName string
	NickName string
	Title string
	Status bool
	Abstract string
	Preferences string
	ContactMe Contact
}

type Data struct {
	Education []Education
	Work []Work
	Skills []Skills
	Services []Services
	Portfolio []Portfolio
	Social []Social
	Personal Personal
}



func main() {
	uri := os.Getenv("MONGOLAB_URL")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	sess, err := mgo.Dial(uri)

	if err != nil {
		log.Fatal("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("homepage").C("user")
	result := Person{}
	err = collection.Find(bson.M{"name": "Prathamesh Sonpatki"}).One(&result)

	/*err = collection.Insert(&Person{"Stefan Klaste", "klaste@posteo.de"},
		&Person{"Nishant Modak", "modak.nishant@gmail.com"},
		&Person{"Prathamesh Sonpatki", "csonpatki@gmail.com"},
		&Person{"murtuza kutub", "murtuzafirst@gmail.com"},
		&Person{"aniket joshi", "joshianiket22@gmail.com"},
		&Person{"Michael de Silva", "michael@mwdesilva.com"},
		&Person{"Alejandro Cespedes Vicente", "cesal_vizar@hotmail.com"})
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

	result := Person{}
	err = collection.Find(bson.M{"name": "Prathamesh Sonpatki"}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}*/

/*
	Education
	Work
	Skills
	Services
	Portfolio
	References
	Contact
	Social
	Personal
*/

	collectionEducation := sess.DB("homepage").C("education")
	collectionWork := sess.DB("homepage").C("work")
	collectionSkills := sess.DB("homepage").C("skills")
	collectionServices := sess.DB("homepage").C("services")
	collectionPortfolio := sess.DB("homepage").C("portfolio")
	collectionSocial := sess.DB("homepage").C("social")
	collectionPersonal := sess.DB("homepage").C("personal")

	var education []Education
	err = collectionEducation.Find(nil).All(&education)

	var work []Work
	err = collectionWork.Find(nil).All(&work)

	var skills []Skills
	err = collectionSkills.Find(nil).All(&skills)

	var services []Services
	err = collectionServices.Find(nil).All(&services)

	var portfolio []Portfolio
	err = collectionPortfolio.Find(nil).All(&portfolio)

	var social []Social
	err = collectionSocial.Find(nil).All(&social)

	var personal Personal
	err = collectionPersonal.Find(nil).One(&personal)

	data := Data{
		education,
		work,
		skills,
		services,
		portfolio,
		social,
		personal}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.Static("/assets", "assets")
	router.Static("/images", "images")



	router.GET("/", func(c *gin.Context) {

		//fmt.Println("Results All: ", data)

		c.HTML(http.StatusOK, "homepage.tmpl.html",
			data)
	})




	router.GET("/mail", func(c *gin.Context) {

		// Set up authentication information.
		auth := smtp.PlainAuth(
			"",
			"example@gmail.com",
			"pass",
			"smtp.gmail.com",
		)
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		err := smtp.SendMail(
			"smtp.gmail.com:25",
			auth,
			"example@gmail.com",
			[]string{"example@gmail.com"},
			[]byte("Mail from Golang."),
		)
		if err != nil {
			log.Fatal(err)
		}else {
			c.JSON(200, gin.H{
				"message": "successful mail send",
			})
		}



	})


	router.GET("/install", func(c *gin.Context) {
		install()
		c.JSON(200, gin.H{
			"message": "successful installation",
		})
	})

	router.Run(":" + port)
}

func install()  {
	uri := os.Getenv("MONGOLAB_URL")

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	sess, err := mgo.Dial(uri)

	if err != nil {
		log.Fatal("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}

	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})


	collectionEducation := sess.DB("homepage").C("education")
	collectionWork := sess.DB("homepage").C("work")
	collectionSkills := sess.DB("homepage").C("skills")
	collectionServices := sess.DB("homepage").C("services")
	collectionPortfolio := sess.DB("homepage").C("portfolio")
	collectionSocial := sess.DB("homepage").C("social")
	collectionPersonal := sess.DB("homepage").C("personal")

	err = collectionEducation.Insert(
		&Education{"unmsm.png","Universidad Nacional de San Marcos", "Software Engineering", "2015 - present"},
		&Education{"cibertec.png","ISTP Cibertec", "Computer's science", "2008 - 2012"},
	)
	err = collectionWork.Insert(
		&Work{"epictrim.png","Epictrim", "Web Developer","Backend", "2014 - 2015"},
		&Work{"apoyo.png","Apoyo Publicitario", "FullStack Developer","TI", "2011 - 2014"},
		&Work{"upc.png","Universidad Peruana de Ciencias Aplicadas", "Trainee","Lirary", "2010 - 2011"},
	)
	err = collectionSkills.Insert(
		&Skills{"php.png","blue","PHP / Zend", 90},
		&Skills{"js.png","blue","JS / React / NodeJS ", 80},
		&Skills{"java.png","blue","Java / Spring", 75},
		&Skills{"ruby.png","blue","Ruby / RoR", 55},
		&Skills{"python.png","blue","Python / Tornado", 45},
		&Skills{"go.png","blue","Go / Gin", 50},
		&Skills{"elixir.png","blue","Elixir / Phoenix", 45},
		&Skills{"net.png","blue","C# / .Net / Mono", 40},
	)

	err = collectionServices.Insert(
		&Services{"http","deep-purple","Web Development", "Lorem"},
		&Services{"android","deep-orange","Mobile Development", "Lorem"},
		&Services{"bug_report","cyan","Bugs", "Lorem"},
	)

	err = collectionPortfolio.Insert(
		&Portfolio{"Dtodoaqui","dtoto_small.png","Web Development", "Sitio Web de Dtodoaqui", "","", ""},
		&Portfolio{"Apoyo Publicitario","apoyo_small.png","Web Development", "Sitio Web de Apoyo Publicitario", "","", ""},
		&Portfolio{"Storewars","storewars_small.png","Web Development", "Sitio Web de Storewars Perú", "","", ""},
		&Portfolio{"MakiDrivePy","makidrive_small.png","","", "","", ""},
	)


	err = collectionSocial.Insert(
		&Social{"facebook","","Facebook"},
		&Social{"twitter","","Twitter"},
		&Social{"google-plus","","Google+"},
		&Social{"github","","Github"},
		&Social{"linkedin","","Linkedin"},
	)

	ContactMe := Contact{"Contact Me","cesar@tineo.mobi","+511 999 666 567","",""}

	err = collectionPersonal.Insert(
		&Personal{"Cesar",
			"Gutierrez","Tineo","Software Developer",true,"",
			"I love <strong>Cats</strong>, <strong>Books!</strong>",
			ContactMe,
		},
	)

	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

	//result := Person{}

	//err = collection.Find(bson.M{"name": "Prathamesh Sonpatki"}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}
}