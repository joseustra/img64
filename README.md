# img64

## local image cache.

**img64** will download and cache images in base64 on you local machine. 

### Instalation

Download the go package

```
go install https://github.com/ustrajunior/img64
```

then start with 

```
PORT="8800" CORS="*" DBNAME="images" img64
```

### Example of use

If the img64 is running on localhost:8800, then it should be used like this:

#### To get the base64 version:

It will the the image passed on the url parameter, download, convert to base64 and save on a local storage (boltdb) then return the base64 content of the image. Notice there is a **/b** before the url parameter.

```
curl "http://localhost:8800/b?url=http://site.com/img.jpg"
```

#### To get the image file:

It will the the image passed on the url parameter, download, convert to base64 and save on a local storage (boltdb) then return the image file.

```
curl "http://localhost:8800?url=http://site.com/img.jpg"
```

When you request the same image url for the second time, it will simply get the image on the storage and return the base64 or image content. So, no network connection will be done to retrieve the image.

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
      $("img").each(function(){                                                                                        
        var  _this = $(this);
        var url = _this.data("image");
        if (url != undefined) {
          $.ajax({
            url: _this.data("image"),      
          }).done(function(data){                                                                                            
            _this.attr("src", "data:image/x;base64," + data);                                                               
          });
        }                                                                                                                                                                                                           
      });
    });                                                                                                                            
  </script>                                                                                                                      
</head>
<body>
  <img src="" style="width:200px" data-image="http://localhost:8800/b?url=http://example.com/image.jpg" alt="image" />
</body>
</html>
```

To use the image file, just do like a normal *<img>* tag with the **http://localhost:8800** as a prefix.

```html
<html>
<head></head>
<body>
  <img src="" style="width:200px" data-image="http://localhost:8800?url=http://example.com/image.jpg" alt="image" />
</body>
</html>
```
## Status
**Not ready for production**
