
Kakao.init('c089c8172def97eb00c07217cae17495')
function kakaoLogout() {
    if (!Kakao.Auth.getAccessToken()) {
        alert('Not logged in.')
        return
    }
    Kakao.Auth.logout(function () {
        alert('logout ok\naccess token -> ' + Kakao.Auth.getAccessToken())
    })
}
