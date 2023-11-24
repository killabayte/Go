package middleware

var (
    SecretKey = []byte("BakerStreet221BNowayToHack")
    emptyValidFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
        return SecretKey, nil
    }
)

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: emptyValidFunc,
    SigningMethod:       jwt.SigningMethodHS256,
})