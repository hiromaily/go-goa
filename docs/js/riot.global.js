(function(){
    var rg = function(){};
    //public
    rg.debug = 1;

    var host = 'http://localhost:8090';

//------------------------------------------------------------------------------
//public
//------------------------------------------------------------------------------
    // set resume page
    rg.setResume = (riot, id) => {
        // data definition
        if (!id) {
            id = sessionStorage.getItem('id')
        }
        let data = [
            {url: `${host}/api/user/${id}/liketech`, element: 'like-tech'},
            {url: `${host}/api/user/${id}/disliketech`, element: 'dislike-tech'},
            {url: `${host}/api/user/${id}/workhistory`, element: 'work-history'}
        ];
        if (window.debugMode == 1){
            data[0].url='/json/liketech.json';
            data[1].url='/json/disliketech.json';
            data[2].url='/json/workhistory.json';
        }


        //loop
        data.forEach((obj) => {
            let fn = (json, element) => {
                if(json.status && json.status != 200){
                    sessionStorage.removeItem('jwt');
                    sessionStorage.removeItem('id');
                    location.href = '/login.html';
                    return;
                }
                //Success
                riot.mount(element, {
                    items : json
                } );
            }

            rg.callAPI(obj.url, {}, 'GET', fn, obj.element)
        });

    };

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