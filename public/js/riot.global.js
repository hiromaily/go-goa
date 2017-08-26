(function(){
    var rg = function(){};
    //public
    rg.debug = 1;

//------------------------------------------------------------------------------
//public
//------------------------------------------------------------------------------

    //common fetch API function
    rg.callAPI = (url, payload, mtd, fn, obj) => {
        let key = sessionStorage.getItem('jwt')
        let body = {
            method: mtd,
            headers: {
                'Authorization': 'Bearer '+key,
                "Content-Type": "application/json"
            }
        }
        if (mtd != "GET" && payload){
            body.body = JSON.stringify(payload)
        }

        fetch(url, body).then((response) => {
            return response.json()
        }).then((json) => {
            console.log("res:", json)
            fn(json, obj)
        });

        return
    }

    window.rg = rg;
})();