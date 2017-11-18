### Blueberry 
(Name to be changed)

Work in progress microservice framework

**Done:**
- Controller handling
- Route resolving
- Events (Pre-Request, Post-Request)
- Dependency injection (into Action)
- Plain Text Responses
- JSON Resposnes

**Work in Progress:**
- Cookies & Sessions

**To be done:**
- Find a name
- Job scheduler

#### FAQ 
Really frequently asked questions from feedback that I got mostly from other developers.

Q: Why so much reflection? <br />
A: Write less code, implement your idea and have it run as soon as possible. I don't want to spend hours writing config
files. I found the conventions of this framework to be suitable for most use cases. Most things can be still configured
manually if the conventions are not suitable. Most reflection and additional complexity is handled in the start-up phase 
of the application. Means, when adding controllers, injectors and events. Later, those get only resolved, 
to keep the runtime complexity low and the performance high (resolves ca. 5263 requests / second on a Raspberry PI (3))

Q: Where are my formatted HTML responses? <br />
A: This is a microservice framework. So HTML responses are not a priority. Go has a nice templating system. 
Connecting it to the response package (By implementing the RenderableResponse interface) would be quite easy,
but does not need to be in the main package. 
Also making standalone (PHP Style) MVC web applications feels a little uncomfortable in Go. 
I'd rather use PHP with Symfony or ASP.NET for such things.

Q: Where is the ORM? <br />
A: There are some nice Go ORMs which have a larger codebase than this framework itself alone. The Beego ORM(1) is nice 
to use and probably the most advanced ORM in Go. Personally, I prefer using SQLX(2), which is just an object mapper
and does not do much magic.


(1) https://beego.me/docs/mvc/model/orm.md
(2) https://github.com/jmoiron/sqlx
(3) 7 Methods, 1 additional injection, 1 additional event. Processing time of additional code not included