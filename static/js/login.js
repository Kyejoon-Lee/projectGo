   Kakao.init('9a11092b4299be7875a3ced245dcfab8')
    Kakao.Auth.createLoginButton({
        container: '#kakao-login-btn',
        success: function(authObj) {
            Kakao.API.request({
                url: '/v2/user/me',
                success: function(res) {
                    alert(JSON.stringify(res))
                },
                fail: function(error) {
                    alert(
                        'login success, but failed to request user information: ' +
                        JSON.stringify(error)
                    )
                },
            })
        },
        fail: function(err) {
            alert('failed to login: ' + JSON.stringify(err))
        },
    })
