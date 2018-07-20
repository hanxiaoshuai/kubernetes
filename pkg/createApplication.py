#coding=utf-8
from selenium import webdriver
import time
#from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.action_chains import ActionChains

driver=webdriver.Chrome()
url = "https://10.XXX:31943/"
driver.get(url)
driver.maximize_window()
driver.implicitly_wait(30)
#driver.find_element_by_id("kw").clear()
driver.find_element_by_id("subUserLogin").click()
time.sleep(3)  
driver.find_element_by_xpath("//li[@id='accountNameLiId']/span/input").send_keys(u"cfe-test")

time.sleep(1)  
driver.find_element_by_xpath("//span[@id='userNameId']/input").send_keys(u"tuser")
time.sleep(1)  

driver.find_element_by_xpath("//span[@id='pwdId']/input").send_keys(u"Huawei@123")
time.sleep(1) 


driver.find_element_by_id("btn_submit").click()
time.sleep(5)


driver.find_element_by_xpath("//div[@class='nui-skin-white nui-variable-light']//span[@title='应用']").click()
time.sleep(5)
driver.find_element_by_xpath("//div[@class='nui-skin-white nui-variable-light']/.././div[6]//span[@title='应用管理']").click()
time.sleep(10)

driver.find_element_by_xpath("//button[@id='appCreateBtnId']").click()
time.sleep(5)

driver.find_element_by_xpath("//h2[contains(text(),'容器应用')]/../div[contains(text(),'创建')]").click()
time.sleep(5)


driver.find_element_by_xpath("//span[contains(text(),'应用名称')]/../../div[2]//input").click()
time.sleep(5)
#driver.find_element_by_xpath("//span[contains(text(),'应用名称')]").send_keys(u"test1")

driver.find_element_by_xpath("//span[contains(text(),'应用名称')]/../../div[2]//input").send_keys(u"test1")
#driver.find_element_by_xpath("//input[@class='tiny-input-text valid_error_input']").send_keys(u"test001")
time.sleep(5)

next=driver.find_element_by_xpath("//button[@class='bigbtn btn btn-primary ng-binding']")
time.sleep(5)

action = ActionChains(driver)
action.move_to_element(next).perform()
time.sleep(10)
next.click()

time.sleep(5)
add1=driver.find_element_by_xpath("//div[@id='appContainerConfig']//i[@class='addimg']")
add1.click()
time.sleep(5)

imgsearch=driver.find_element_by_xpath("//input[@class='tiny-searchBox-input']")
imgsearch.send_keys(u"nginx")
time.sleep(5)

search=driver.find_element_by_xpath("//div[@class='ti ti-search tiny-searchBox-searchIcon']")
search.click()
time.sleep(5)

driver.find_element_by_xpath("//td[@tdtitle='镜像仓库名称']/../td[1]//div[@class='tiny-radioIcon tiny-radio-unchecked']").click()
time.sleep(5)
driver.find_element_by_xpath("//div[@class='ui-dialog-buttonpane ui-widget-content ui-helper-clearfix']//span[contains(text(),'确认')]").click()

time.sleep(5)

save=driver.find_element_by_xpath("//button[@class='smallbtn bluebtn ng-binding ng-scope']")
time.sleep(5)

action = ActionChains(driver)
action.move_to_element(save).perform()
time.sleep(5)
save.click()
time.sleep(5)


next2=driver.find_element_by_xpath("//button[@class='bigbtn btn btn-primary ng-binding']")
next2.click()
time.sleep(5)

nnoo=driver.find_element_by_xpath("//div[@class='tiny-radioIcon tiny-radio-unchecked']")
nnoo.click()
time.sleep(5)


driver.find_element_by_xpath("//button[@class='bigbtn btn btn-primary ng-binding']").click()



ccreate=driver.find_element_by_xpath("//button[@class='bigbtn btn btn-primary ng-binding']")
action = ActionChains(driver)
action.move_to_element(ccreate).perform()
time.sleep(5)
ccreate.click()

#driver.find_element_by_xpath("//div[@class='nui-skin-white nui-variable-light']//span[@title='编排']").click()
#time.sleep(2) 

#driver.find_element_by_xpath("//div[@class='nui-skin-white nui-variable-light']/../div[6]//span[@title='模板']").click()
#time.sleep(3)
#driver.find_element_by_xpath("//button[@class='smallbtn whitebtn addItemBtn ng-binding ng-scope']").click()
#time.sleep(3)

#driver.find_element_by_xpath("//label[@class='color-blue underline pointerClass margintop-15 ng-binding']").click()
#time.sleep(3)

#driver.find_element_by_xpath("//div[@id='aosPackage']//td[contains(text(),'test002')]/../td[1]//tbody//td[1]").click()
#time.sleep(3)
#driver.find_element_by_xpath("//div[@class='ui-dialog-buttonpane ui-widget-content ui-helper-clearfix']//span[contains(text(),'确定')]/..").click()
#time.sleep(3)




#driver.find_element_by_xpath("//span[contains(text(),'模板名称')]/../..//input[@placeholder='以小写字母、数字或者中划线（-）三种字符中一种或一种以上组合，必须以小写字母开头，小写字母或者数字结尾，输入字符在24个以内']").send_keys(u"irvingtest")


#time.sleep(3)
#driver.find_element_by_xpath("//span[contains(text(),'版本')]/../..//input").send_keys(u"v0.9")

#time.sleep(3)

#driver.find_element_by_xpath("//div[@class='handle-list clearfix']//span[contains(text(),'创建')]").click()
#time.sleep(6)


#driver.find_element_by_xpath("//span[contains(text(),'irvingtest')]/../..//i[@tooltip='删除']").click()
#time.sleep(5)




#driver.find_element_by_xpath("//div[@class='ui-dialog-buttonpane ui-widget-content ui-helper-clearfix']//span[contains(text(),'确定')]").click()

#time.sleep(5)

#time.sleep(10)   




