package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/yourusername/ai-backend/controllers"
    "github.com/yourusername/ai-backend/models"
    "github.com/yourusername/ai-backend/config"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // 用户注册路由
    r.POST("/register", func(c *gin.Context) {
        var user models.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input"})
            return
        }

        // 创建用户
        newUser, err := user.CreateUser(config.DB)
        if err != nil {
            c.JSON(400, gin.H{"error": "Could not create user"})
            return
        }

        // 生成 JWT
        token, err := controllers.GenerateToken(*newUser)
        if err != nil {
            c.JSON(400, gin.H{"error": "Could not generate token"})
            return
        }

        c.JSON(200, gin.H{"token": token})
    })

    // 用户登录路由
    r.POST("/login", func(c *gin.Context) {
        var user models.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input"})
            return
        }

        // 模拟登录验证
        var dbUser models.User
        if err := config.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
            c.JSON(400, gin.H{"error": "Invalid username or password"})
            return
        }

        // 生成 JWT
        token, err := controllers.GenerateToken(dbUser)
        if err != nil {
            c.JSON(400, gin.H{"error": "Could not generate token"})
            return
        }

        c.JSON(200, gin.H{"token": token})
    })

    // 受保护的路由
    protected := r.Group("/protected")
    protected.Use(controllers.AuthMiddleware())
    {
        protected.GET("/data", func(c *gin.Context) {
            c.JSON(200, gin.H{"message": "This is protected data"})
        })
    }

    return r
}