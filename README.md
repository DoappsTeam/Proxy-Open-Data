#Hackathon San Isidro
This is a proxy for wrap the request towards the base API.

##Environment
YOUR_API_KEY:
41bcaadbfc3362545c067bdf0466b01bc247c84d

GUIDs:
JORNA-Y-CAMPA-SALUD
ACTIV-GRATU
CAMPA-VETER
ACTIV-ATENC-A-LA-PERSO
HORAR-Y-RUTAS-RECOL-RESID
CHAPA-TU-BICI-CICLO-EXIST
ACTIV-CULTU-86315
PRESU-INGRE-2015-57742
PRESU-GASTO-2015
CONSO-INTER-SEREN-31211

PATH:
http://api.datosabiertos.msi.gob.pe/datastreams/invoke/GUID?auth_key=YOUR_API_KEY

#Internals
```go
    // First al all, the data is classified
        objTest := `
        {
            "id": 10,
            "name": "cultura",
            "result": {
                "cols": 2,
                "rows": 2,
                "length": 3,
                "array": [
                    { "data": "ID" },
                    { "data": "Name" },
                    { "data": "1" },
                    { "data": "Calle" },
                    { "data": "2" },
                    { "data": "Avenida" }
                ]
            }
        }
        `
    // To    
    
        [
            {
                "ID": "1",
                "Name": "Calle"
            },
            {
                "ID": "2",
                "Name": "Avenida"
            }
        ]
    
```


#Usage
path ´/api/:namespace´

#TODO :rocket:
- Documentation
- Define paths