## 接单管理系统 API 文档

这是一个基于 Go 语言开发的接单管理系统，使用 SQLite3 作为数据库。

---

### **Admin API**

#### **1. 登录**

- **URL**: `/api/v1/auth/login`
- **Method**: `POST`
- **Request Body** (JSON):
    ```json
    {
        "username": "string",
        "password": "string",
        "captcha": "string"
    }
    ```

---

#### **2. 创建管理员**

- **URL**: `/api/v1/admin`
- **Method**: `POST`
- **Request Body** (JSON):
    ```json
    {
        "username": "string",
        "password": "string",
        "role": "string",
        "profile": {
            "name": "string",
            "phone": "string",
            "email": "string"
        }
    }
    ```

---

#### **3. 获取管理员列表**

- **URL**: `/api/v1/admin`
- **Method**: `GET`
- **Query Parameters**:
    - `page`: `int` (可选，默认值: `1`)
    - `size`: `int` (可选，默认值: `10`)
    - `role`: `string` (可选)

---

#### **4. 获取管理员个人信息**

- **URL**: `/api/v1/admin/:id/profile`
- **Method**: `GET`

---

#### **5. 更新管理员信息**

- **URL**: `/api/v1/admin/:id`
- **Method**: `PUT`
- **Request Body** (JSON):
    ```json
    {
        "role": "string",
        "profile": {
            "name": "string",
            "phone": "string",
            "qq": "string",
            "wx": "string",
            "email": "string"
        }
    }
    ```

---

#### **6. 删除管理员**

- **URL**: `/api/v1/admin/:id`
- **Method**: `DELETE`

---

### **Demand API**

#### **1. 创建需求**

- **URL**: `/api/v1/demand`
- **Method**: `POST`
- **Request Body** (JSON):
    ```json
    {
        "title": "string",
        "descs": "string",
        "images": "string",
        "attachments": "string",
        "price": "string",
        "status": "string",
        "user_id": "integer"
    }
    ```

---

#### **2. 获取需求**

- **URL**: `/api/v1/demand/:id`
- **Method**: `GET`

---
#### **3. 获取全部需求列表**

- **URL**: `/api/v1/demand`
- **Method**: `GET`

---

#### **4. 更新需求**

- **URL**: `/api/v1/demand/:id`
- **Method**: `PUT`
- **Request Body** (JSON):
    ```json
    {
        "title": "string",
        "descs": "string",
        "images": "string",
        "attachments": "string",
        "price": "string",
        "status": "string",
        "user_id": "integer"
    }
    ```

---

#### **5. 删除需求**

- **URL**: `/api/v1/demand/:id`
- **Method**: `DELETE`

---

### **Customer API**

#### **1. 创建用户**

- **URL**: `/api/v1/customer`
- **Method**: `POST`
- **Request Body** (JSON):
    ```json
    {
        "name": "string",
        "phone": "string",
        "qq": "string",
        "wx": "string",
        "total": "string"
    }
    ```

---

#### **2. 获取用户**

- **URL**: `/api/v1/customer/:id`
- **Method**: `GET`

---

#### **3. 获取所有用户**

- **URL**: `/api/v1/customer/`
- **Method**: `GET`

---

#### **4. 更新用户**

- **URL**: `/api/v1/customer/:id`
- **Method**: `PUT`
- **Request Body** (JSON):
    ```json
    {
        "name": "string",
        "phone": "string",
        "qq": "string",
        "wx": "string",
        "total": "string"
    }
    ```

---

#### **5. 删除用户**

- **URL**: `/api/v1/customer/:id`
- **Method**: `DELETE`