### Данные приложения

при компиляции добавьте 
$_PKG_ - путь к переменным

_PKG_="bot/pkg/app/build"

```
 -ldflags="-X '$_PKG_.version=$BUILD_VERSION' -X '$_PKG_.hashCommit=$BUILD_HASH_COMMIT' -X '$_PKG_.branch=$BUILD_BRANCH' -X '$_PKG_.time=$BUILD_TIME'"
```


Для разработки можно установить переменные окружения: 
Значения, установленные при компиляции, имеют больший приоритет
```
ENV_PKG_APP_BUILD_DEV_MODE_VERSION=
ENV_PKG_APP_BUILD_DEV_MODE_HASH_COMMIT=
ENV_PKG_APP_BUILD_DEV_MODE_BRANCH=
ENV_PKG_APP_BUILD_DEV_MODE_TIME=

```
