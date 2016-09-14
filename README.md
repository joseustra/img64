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

```
curl "http://localhost:8800?url=http://site.com/img.jpg"
```

It will the the image passed on the url parameter, download, convert to base64 and save on a local storage (boltdb) then return the base64 content of the image.

When you request the same image url for the second time, it will simply get the image on the storage and return the base64 content. So, no network connection will be done to retrieve the image.

Let's see a simple example on a webpage.

```html
<html>
<head>
  <script src="jquery.js"></script>
  <script type="text/javascript" charset="utf-8">                                
  $(document).ready(function() {    
    $("img").each(function(){                                                                                        
      var  _this = $(this);                                                                                               
      $.ajax({                                                            
        
        url: _this.data("image"),      
      }).done(function(data){                                                                                            
        _this.attr("src", "data:image/x;base64," + data);                                                               
      });                                                                                                                 
    });
  });                                                                                                                            
  </script>                                                                                                                     
</head>
<body>
  <img src="" data-image="http://url.com/image.jpg" alt="kitesurfing" />                                                                                
</body>
</html>
```

## Status
**Not ready for production**
