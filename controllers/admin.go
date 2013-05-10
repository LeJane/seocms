package controllers

import (
    //"github.com/astaxie/beedb"
    "github.com/astaxie/beego"
)

type AdminController struct {
    beego.Controller
}

var err error

func (this *AdminController) Get() {
    this.Layout = "layout_admin.tpl"
    object := this.Ctx.Params[":object"]
    action := this.Ctx.Params[":action"]
    if object == "article" {
        switch action {
        case "list":
            this.Data["PageTitle"] = "文章列表_文章管理_SEOCMS"
            this.Data["Articles"] = []string{
                "这是第1篇博客",
                "这是第2篇博客",
                "这是第3篇博客",
            }
            this.TplNames = "admin/article_list.tpl"
        case "add":
            this.Data["PageTitle"] = "添加文章_文章管理_SEOCMS"
            this.Data["Categories"] = []string{"博客", "笔记"}
            this.TplNames = "admin/add_article.tpl"
        case "edit":
            this.Data["Id"] = this.Ctx.Params[":id"]
            this.Data["PageTitle"] = "编辑文章_文章管理_SEOCMS"
            this.TplNames = "admin/edit_article.tpl"
        case "delete":
            this.Data["Id"] = this.Ctx.Params[":id"]
            this.Data["PageTitle"] = "删除文章_文章管理_SEOCMS"
            this.TplNames = "admin/del_article.tpl"
        }
    } else if object == "category" {
        switch action {
        case "list":    // 分类列表
            categories := []Category{}
            orm := InitDb()
            err = orm.OrderBy("name").FindAll(&categories)
            Check(err)
            this.Data["PageTitle"] = "分类列表_文章管理_SEOCMS"
            this.Data["Categories"] = categories
            this.TplNames = "admin/category_list.tpl"
        case "add":    // 添加分类
            this.Data["PageTitle"] = "添加分类_文章管理_SEOCMS"
            this.TplNames = "admin/add_category.tpl"
        case "edit":    // 修改分类
            id := this.Ctx.Params[":id"]
            //this.Data["Id"] = id

            orm := InitDb()
            category := Category{}
            err = orm.Where("id=?", id).Find(&category)
            Check(err)
            this.Data["Category"] = category

            this.Data["PageTitle"] = "修改分类_文章管理_SEOCMS"
            this.TplNames = "admin/edit_category.tpl"
        case "delete":    // 删除分类
            //this.Data["Id"] = this.Ctx.Params[":id"]
            //this.Data["PageTitle"] = "删除分类_文章管理_SEOCMS"
            //this.TplNames = "admin/delete_category.tpl"
            id := this.Ctx.Params[":id"]

            orm := InitDb()
            category := Category{}
            err = orm.Where("id=?", id).Find(&category)
            Check(err)
            orm.Delete(&category)

            this.Ctx.Redirect(301, "/category/list")
        }
    }
}

func (this *AdminController) Post() {
    object := this.Ctx.Params[":object"]
    action := this.Ctx.Params[":action"]
    if object == "article" {
        this.Ctx.Redirect(302, "/article/list")
    } else if object == "category" {
        switch action {
        case "add":    // 处理添加分类
            name := this.Input().Get("name")
            nameEn := this.Input().Get("name_en")
            description := this.Input().Get("description")
            alias := this.Input().Get("alias")

            orm := InitDb()
            category := Category{}
            err = orm.Where("name=?", name).Find(&category)
            if err != nil {
                orm = InitDb()
            }

            category.Name = name
            category.NameEn = nameEn
            category.Description = description
            category.Alias = alias
            //category.Alias = Str2slice(alias)
            //Debug("%s, %s, %s, %v", name, nameEn, description, alias)
            Debug("%s, %s, %s, %v", category.Name, category.NameEn, category.Description, category.Alias)
            err = orm.Save(&category)
            Check(err)

            this.Ctx.Redirect(302, "/category/list")
        case "edit":    // 处理修改分类
            id := this.Ctx.Params[":id"]

            orm := InitDb()
            category := Category{}
            err = orm.Where("id=?", id).Find(&category)
            Check(err)

            category.Name = this.Input().Get("name")
            category.NameEn = this.Input().Get("name_en")
            category.Description = this.Input().Get("description")
            category.Alias = this.Input().Get("alias")
            err = orm.Save(&category)
            Check(err)

            this.Ctx.Redirect(302, "/category/list")
        }
    }
}
