package handler

import (
    "morag_ecommerce/entity"
    "net/http"
    "time"
    "github.com/dgrijalva/jwt-go"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
)

// Buat fungsi createJWTToken
func createJWTToken(userID uint) (string, error) {
    // Atur claims
    claims := jwt.MapClaims{}
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 2).Unix() 

    // Buat token dengan tanda tangan
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    
    secret := []byte("YourSecretKey")

    
    tokenString, err := token.SignedString(secret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}



func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// Untuk memeriksa password yang diinputkan dengan hashed password yang tersimpan di database
func verifyPassword(inputPassword, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
    return err == nil
}

// Item godoc
// @Summary Register
// @Schemes
// @Description Register
// @Accept json
// @Produce json
// @Param store body entity.Store true "Register New User"
// @Success 201 {object} entity.Store
// @Router /Register [post]
// @Tags data user
func RegisterStore(c *gin.Context) {
    var store entity.Store
    if err := c.ShouldBindJSON(&store); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if store.StoreEmail == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "store_email is required"})
        return
    }

    if len(store.Password) < 8 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
        return
    }

    if len(store.StoreName) < 6 || len(store.StoreName) > 15 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "store_name must be between 6 and 15 characters"})
        return
    }

    if store.StoreType == "" {
        store.StoreType = "silver" 
    }

    hashedPassword, err := hashPassword(store.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    store.Password = hashedPassword 

    db := c.MustGet("db").(*gorm.DB)
    if err := db.Create(&store).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save store data"})
        return
    }

    token, err := createJWTToken(store.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "store_id":    store.ID,
        "store_email": store.StoreEmail,
        "store_name":  store.StoreName,
        "store_type":  store.StoreType,
        "token":       token,
    })
}

func LoginStore(c *gin.Context) {
    var Store entity.Store
    if err := c.ShouldBindJSON(&Store); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if Store.StoreEmail == "" || Store.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Store email and password are required"})
        return
    }

    db := c.MustGet("db").(*gorm.DB)
    var store entity.Store

    if err := db.Where("store_email = ?", Store.StoreEmail).First(&store).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find the store"})
        }
        return
    }

    if !verifyPassword(Store.Password, store.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
        return
    }

    token, err := createJWTToken(store.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "store_id":    store.ID,
        "store_email": store.StoreEmail,
        "store_name":  store.StoreName,
        "store_type":  store.StoreType,
        "token":       token,
    })
}
