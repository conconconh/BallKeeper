<html>
<head>
<title></title>
</head>
<body>
<form action="/loginSubmit" method="post">
    使用者名稱:<input type="text" name="username">
    密碼:<input type="password" name="password">
    <input type="submit" value="提交">
</form>

<form action="/changePassword" method="put">
    使用者名稱:<input type="text" name="username">
    新密碼:<input type="password" name="password">
    <input type="submit" value="更新">
</form>

<form action="/deleteUser" method="depete">
    欲刪除使用者名稱:<input type="text" name="username">
    <input type="submit" value="刪除">
</form>

</body>
</html>