[1mdiff --git a/devfile.yaml b/devfile.yaml[m
[1mindex 799d921..9940767 100644[m
[1m--- a/devfile.yaml[m
[1m+++ b/devfile.yaml[m
[36m@@ -1,6 +1,6 @@[m
 schemaVersion: 2.2.2[m
 metadata:[m
[31m-  name: golang[m
[32m+[m[32m  name: oaimg[m
 components:[m
   - name: tools[m
     container:[m
[36m@@ -17,7 +17,7 @@[m [mcomponents:[m
           value: /tmp/.cache[m
       endpoints:[m
         - exposure: public[m
[31m-          name: 'health-check'[m
[32m+[m[32m          name: web[m
           protocol: https[m
           targetPort: 8080[m
 commands:[m
