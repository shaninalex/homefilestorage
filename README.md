## HomeFileStorage

## Note!

It's my infinite pet-project. Here I'm not actualy developming filestorage... This is sandbox where I test different approaches in software development. And filestorage is just super complex todo app where I just solving some common tasks - access management, performance, comunications between services and so on... So here I can make some mistakes. If you realy need some self hosted filestorage look at Nextcloud or ownCloud. Or if you curious person as I'm - please feel free to contribute!


## Now

*Now I'm trying to implement some authentication and authorization services instead of making them from screach. Because there are so many existed solutions for this! And making "yeat another user service" makes me sad. What the point of pure tested and pure secured user/auth services? It will be better to integrate Ory services, Keycloack or Zitadel and deligate them user management... And create upload large files ( streaming ) functionality instead.*

## Docs

Developer documentation and tasks will be soon.

### Start

```bash
$ make restart
```

## Front end app

[Homefilestorage frontend](https://github.com/shaninalex/homefilestorage-frontend)


```bash
$ npm run start
```

Make sure that IP (or ULR ) in proxy.conf.json is the same as oathkeeper
