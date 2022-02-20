# fakedouban

#所实现的接口
#用户相关：     
登录 http://119.91.20.70:6060/user/login  
注册 http://119.91.20.70:6060/user/Singup  
修改密码 http://119.91.20.70:6060/user/Reset   
获取密保问题 http://119.91.20.70:6060/user/QueryProtectionQ  
修改个人简介 http://119.91.20.70:6060/user/change  
返回登录用户的信息 http://119.91.20.70:6060/user/imfor  
获取其他用户的信息 http://119.91.20.70:6060/user/otherimfor

#电影，影人信息相关：  
查询单个电影信息 http://119.91.20.70:6060/object  
获取主页推荐电影信息 http://119.91.20.70:6060/recommend  
获取单个影人信息 http://119.91.20.70:6060/celebrity  
新片榜单 http://119.91.20.70:6060/newhotlist  
分类热门电影榜单 http://119.91.20.70:6060/classhot  
分类查询电影 http://119.91.20.70:6060/class  
按名字模糊查询电影，人物信息 http://119.91.20.70:6060/search

#讨论区相关：
获取单个讨论区信息 http://119.91.20.70:6060/talking  
发布一个讨论 http://119.91.20.70:6060/Settalking  
在一个讨论下发布评论 http://119.91.20.70:6060/talkingcm  
在一个讨论的评论下发布子评论 http://119.91.20.70:6060/talkingchcm  

#影评与短评相关：
发布一个影评 http://119.91.20.70:6060/comment/parent  
在一个影评下评论 http://119.91.20.70:6060/comment/child  
发布一个短评 http://119.91.20.70:6060/comment/shortcomment  
一个电影下按时间排序的短评 http://119.91.20.70:6060/shortbytime  
一个电影下按有用数排序的短评 http://119.91.20.70:6060/shortbyuse  
一个电影下按有用数排序的评论 http://119.91.20.70:6060/commentbyuse  
一个电影下按时间排序的评论 http://119.91.20.70:6060/commentbytime  
为短评增加有用数或无用数 http://119.91.20.70:6060/comment/scmuse  
  
  
用户登录认证的方式：jwt