# goutil说明

​	简要说明goutil的依赖关系

## 调用关系图

```mermaid
graph BT

A[stringUtil]--> B[logUtil]
A -->  C[fileUtil]
C -->  B

D[idUtil] --> B 
```

