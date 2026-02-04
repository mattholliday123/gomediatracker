What is it
----------------------------------------------------------------
gomediatracker is a web app to track your phsyical media collection.
I use server made in Go that serves the HTML pages and handles http requests. The frontend is very minimal just using HTML, CSS, and vanilla JS. The JS is used mostly to communicate with the server to retrieve data
and render the data on the html pages.

**Progress so far and plans:**

It functions minimally with video games. You are able to search and the API will return top results. You can add them to your collection as well set a status(played or unplayed)
Future plans include allowing to change status of game, clean up frontend, integrate APIs for other physical media such as movies, books, music, etc.
