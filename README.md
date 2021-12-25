# Interactive Sudoku Generator 
This is a polyglot programming project where I tried to implement an unique Sudoku instance generator using 3 different programming languages (Python, Go, Javascript). Containerized using [Docker](https://www.docker.com/).

The project included 2 components:
* Backend which includes the Sudoku instance generating logic written in Go (found in `backend/sudoku/generator/`), and the REST API implemented using [Django Rest Framework](https://www.django-rest-framework.org/). Django is able to called the Go generator function via Python [ctypes library](https://docs.python.org/3/library/ctypes.html).
* Front-end which includes a simple React website (created via [CRA](https://github.com/facebook/create-react-app)) that communicates with the API to grab the generated Sudoku instance and render it. 

### How to run
First create a `.env` file for Django in `backend/backend/`. Include those 2 options: 
```
DEBUG=True
SECRET_KEY=YOUR_SECRET_KEY
```
[See here for how to generate a Django Secret Key.](https://humberto.io/blog/tldr-generate-django-secret-key/)

Then just simply do:
```
docker-compose up
```
After that, navigate to `localhost:3000` in your browser.

### Remarks
The main and only goal I set out when starting on this project is to implement a system with at least 3 different languages, thus learning new technologies (I did not have any experience with Go, Django, React nor Docker prior). Thus, the architecture decisions I made throughout the project were very questionable :sweat_smile:. If someone is to ask to implement the whole thing again without the polyglot requirement, I would have done it all in JS.

Though, it was still a fun project to work on and I have certainly learned a lot from it.
