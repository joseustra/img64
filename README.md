# img64

## local image cache.

**img64** will download and cache images in base64 on you local machine. 

### Instalation

Download the go package

```bash
go get github.com/ustrajunior/img64
```

then start with the port and CORS url

```bash
PORT="8800" CORS="*" DBNAME="images" img64
```

### Example of use

If the img64 is running on localhost:8800, then it should be used like this:

#### To get the base64 version:

It will get the image passed on the url parameter, download, convert to base64 and save on a local storage (boltdb) then return the base64 content of the image. Notice there is a **/b** before the url parameter.

```bash
curl "http://localhost:8800/b?url=http://site.com/img.jpg"
```

#### To get the image file:

It will get the image passed on the url parameter, download, convert to base64 and save on a local storage (boltdb) then return the image file.

```bash
curl "http://localhost:8800?url=http://site.com/img.jpg"
```

When you request the same image url for the second time, it will simply get the image on the storage and return the base64 or image content. So, no network connection, to the original image url, will be done to retrieve the image.

Let's see a simple example on a webpage using the base64 version.

```html
<html>
<head>
  <script
  src="https://code.jquery.com/jquery-3.1.1.min.js"
  integrity="sha256-hVVnYaiADRTO2PzUGmuLJr8BLUSjGIZsDYGmIJLv2b8="
  crossorigin="anonymous"></script>
  <script type="text/javascript" charset="utf-8">                                
		$(document).ready(function() {    
			var _this = $(this);
			var img64 = $(".base64-image");
			var url = img64.data("image");

			if (url != undefined) {
				$.ajax({
					url: url,      
				}).done(function(data){
					img64.attr("src", "data:image/x;base64," + data);
				});
			}
		});
  </script>                                                                                                                      
</head>
<body>
  <img src="" style="width:200px" data-image="http://localhost:8800/b?url=http://example.com/image.jpg" alt="image" />
</body>
</html>
```

To use the image file, just do like a normal *<img>* tag with the **http://localhost:8800?url=** as a prefix.

```html
<html>
<head></head>
<body>
    <img class="image-file" src="http://localhost:8800?url=https://images.unsplash.com/photo-1516843751971-41c5eaaa96fa?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=cd274d78ffb60a5b1e7734adc29b26c3&auto=format&fit=crop&w=1267&q=80" style="width:200px" alt="image" />                                                                          
</body>
</html>
```

## Status
**Not ready for production**

