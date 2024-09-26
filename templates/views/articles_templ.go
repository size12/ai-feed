// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/google/uuid"
)

type Keyword struct {
	Name  string
	Count string
}

type Article struct {
	ID uuid.UUID

	Title       string
	ImageBase64 string
	Content     string

	WordsCount   string
	SymbolsCount string

	Keywords []*Keyword
}

func NewArticles(articles []*Article, themes []*Theme, personalities []*Personality) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"ru\"><head><meta charset=\"UTF-8\"><title>Панель управления</title><link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css\"><script>\n            function validateForm(formData) {\n               for (let [key, val] of formData.entries()) {\n                  if (val === \"\") {\n                     alert(`Поле '${key}' не должно быть пустым`)\n                     return false\n                  }\n               }\n\n               return true\n            }\n\n            function createDialog(id, title = \"\", content = \"\", image = \"\") {\n               const dialogHTML = `\n                             <article>\n                                <header>\n                                   <button aria-label=\"close\" rel=\"prev\" id=\"close-dialog\"></button>\n                                   <p>\n                                      <strong>Редактирование статьи</strong>\n                                   </p>\n                                </header>\n                                <p>\n                                   <u>Измените нужную информацию</u>\n                                </p>\n                                <form id=\"create-form\">\n                                   <label>\n                                   Заголовок статьи\n                                   <input type=\"text\" name=\"title\" placeholder=\"Кликбейтный заголовок 1\" value=\"${title}\">\n                                   </label>\n                                   <label>\n                                   Текст статьи\n                                   <textarea id=\"article-content\" name=\"content\"\n                                      rows=\"10\"\n                                      placeholder=\"Напишите или сгенерируйте статью\"\n                                      aria-label=\"Напишите или сгенерируйте статью\">${content}</textarea>\n                                   </label>\n                                   <label>\n                                   Выберите рисунок\n                                   <input type=\"file\" id=\"image-select\">\n                                   </label>\n                                </form>\n                                <footer>\n                                   <button id=\"dialog-btn\">Сохранить</button>\n                                </footer>\n                             </article>\n                             `\n\n               const dialog = document.createElement('dialog');\n               dialog.open = true;\n               dialog.innerHTML = dialogHTML;\n\n               dialog.querySelector(\"#close-dialog\").onclick = function() {\n                  dialog.parentNode.removeChild(dialog);\n               }\n\n               const form = dialog.querySelector(\"form\")\n\n               var imageContent = \"\"\n\n               dialog.querySelector(\"#image-select\").addEventListener(\"change\", (e) => {\n                  const imageFile = e.target.files[0];\n\n                  var fileReader = new FileReader()\n\n                  fileReader.onload = function(fileLoadedEvent) {\n                     imageContent = fileLoadedEvent.target.result\n                     console.log(imageContent)\n                  }\n\n                  fileReader.readAsDataURL(imageFile);\n               })\n\n               dialog.querySelector(\"#dialog-btn\").addEventListener(\"click\", function(event) {\n                  event.preventDefault()\n\n                  const formData = new FormData(form);\n\n                  if (!validateForm(formData)) {\n                     return\n                  }\n\n                  let obj = Object.fromEntries(formData)\n                  obj[\"id\"] = id\n                  obj[\"image_base64\"] = imageContent\n\n                  const json = JSON.stringify(obj);\n\n                  fetch('/api/article', {\n                     method: \"PUT\",\n                     headers: {\n                        'Content-Type': 'application/json'\n                     },\n                     body: json\n                  }).then(response => {\n                     if (response.status !== 200 && response.status !== 201 && response.status !== 202) {\n                        response.text().then(function(text) {\n                           alert(`Не удалось изменить текст: ${text}`)\n                        })\n                        return\n                     }\n\n                     location.reload();\n\n                  }).catch(err => alert(`Не получилось отправить запрос на сервер ${err}`))\n               })\n\n               document.body.appendChild(dialog)\n            }\n\n            function editArticle(id) {\n               fetch(`/api/article?id=${id}`, {\n                     method: \"GET\",\n                  })\n                  .then((response) => {\n                     response.json().then((article) => {\n                        createDialog(id, article[\"title\"], article[\"content\"])\n                     })\n                  })\n                  .catch(err => alert(`Не получилось отправить запрос на сервер ${err}`))\n            }\n\n            window.onload = function() {\n               const dialogElement = document.querySelector(\"#create-dialog\")\n               const dialogForm = dialogElement.querySelector(\"form\")\n               const themeSelect = dialogElement.querySelector(\"#select-theme\")\n               const personalitySelect = dialogElement.querySelector(\"#select-personality\")\n               const title = dialogElement.querySelector(\"#article-title\")\n               const content = dialogElement.querySelector(\"#article-content\")\n               const articleImage = dialogElement.querySelector(\"#article-image\")\n               const saveArticleBtn = dialogElement.querySelector(\"#save-article-btn\")\n               const metricsBtn = dialogElement.querySelector(\".metrics\")\n\n               document.querySelector(\"#create-article\").onclick = function() {\n                  dialogElement.open = true;\n               }\n\n               dialogElement.querySelector(\"#close-dialog\").onclick = function() {\n                  dialogElement.open = false;\n               }\n\n               dialogElement.querySelector(\"#generate-content\").onclick = function(e) {\n                  e.preventDefault()\n\n                  const themeDescription = themeSelect.options[themeSelect.selectedIndex].text\n\n                  if (themeSelect.selectedIndex === 0) {\n                     alert(\"Пожалуйста выберите тему из списка\")\n                     return\n                  }\n\n                  if (personalitySelect.selectedIndex === 0) {\n                     alert(\"Пожалуйста выберите личность из списка\")\n                     return\n                  }\n\n                  const personality = {\n                     \"name\": personalitySelect.options[personalitySelect.selectedIndex].text,\n                     \"biography\": personalitySelect.options[personalitySelect.selectedIndex].getAttribute(\"data-bio\"),\n                     \"keywords\": personalitySelect.options[personalitySelect.selectedIndex].getAttribute(\"data-keywords\"),\n                     \"thematics\": personalitySelect.options[personalitySelect.selectedIndex].getAttribute(\"data-thematics\"),\n                     \"text_style\": personalitySelect.options[personalitySelect.selectedIndex].getAttribute(\"data-text-style\")\n                  }\n\n\n                  const request = {\n                     \"theme\": {\n                        \"description\": themeDescription\n                     },\n\n                     \"personality\": personality\n                  }\n\n\n                  this.setAttribute(\"aria-busy\", true)\n\n                  fetch('/api/generate/article', {\n                        method: 'POST',\n                        headers: {\n                           'Content-Type': 'application/json'\n                        },\n                        body: JSON.stringify(request)\n                     }).then(response => {\n                        this.setAttribute(\"aria-busy\", false)\n                        if (response.status !== 200) {\n                           response.text().then(function(text) {\n                              console.log(text)\n                              alert(`Не получилось сгенерировать статью: ${text}`)\n                           })\n\n                           return\n                        }\n                        response.json().then(function(response) {\n                           title.value = response[\"title\"]\n                           content.value = response[\"content\"]\n                        })\n\n                     }).catch(err => alert(`Не получилось отправить запрос на сервер ${err}`))\n                     .finally(() => this.setAttribute(\"aria-busy\", false))\n               }\n\n               document.querySelectorAll(\".article\").forEach(el => {\n                  const id = el.getAttribute(\"data-uuid\")\n\n                  const description = el.querySelector(\"h5\").textContent\n\n                  el.querySelector(\"button.open\").onclick = () => {\n                     window.open(`/article/${id}`, '_blank').focus();\n                  }\n\n                  el.querySelector(\"button.edit\").addEventListener(\"click\", function() {\n                     editArticle(id)\n                  })\n\n                  el.querySelector(\"button.delete\").addEventListener(\"click\", function() {\n                     fetch('/api/article', {\n                        method: 'DELETE',\n                        headers: {\n                           'Content-Type': 'text/plain'\n                        },\n                        body: id\n                     }).then(response => {\n                        if (response.status !== 200) {\n                           response.text().then(function(text) {\n                              alert(`Не получилось удалить статью: ${text}`)\n                           })\n                           return\n                        }\n\n                        location.reload();\n\n                     }).catch(err => alert(`Не получилось отправить запрос на сервер ${err}`))\n                  })\n               })\n\n               const generateImageBtn = document.querySelector(\"#generate-image\")\n\n               generateImageBtn.onclick = (e) => {\n                  e.preventDefault()\n\n                  if (!title.value) {\n                     alert(\"Нет заголовка, не можем сгенерировать картинку\")\n                     return\n                  }\n\n                  if (!content.value) {\n                     alert(\"Статья пустая, не можем сгенерировать картинку\")\n                     return\n                  }\n\n                  const request = {\n                     \"title\": title.value,\n                     \"content\": content.value\n                  }\n\n                  generateImageBtn.setAttribute(\"aria-busy\", true)\n\n                  fetch('/api/generate/image', {\n                        method: 'POST',\n                        headers: {\n                           'Content-Type': 'application/json'\n                        },\n                        body: JSON.stringify(request)\n                     }).then(response => {\n                        if (response.status !== 200) {\n                           response.text().then(function(text) {\n                              alert(`Не получилось сгенерировать картинку: ${text}`)\n                           })\n                           return\n                        }\n\n                        response.json().then(function(response) {\n                           articleImage.src = response[\"image_base64\"]\n                        })\n\n\n                     }).catch(err => alert(`Не получилось отправить запрос на сервер ${err}`))\n                     .finally(() => generateImageBtn.setAttribute(\"aria-busy\", false))\n               }\n\n               saveArticleBtn.onclick = () => {\n                  if (!title.value) {\n                     alert(\"Нет заголовка\")\n                     return\n                  }\n\n                  if (!content.value) {\n                     alert(\"Статья пустая\")\n                     return\n                  }\n\n                  const request = {\n                     \"title\": title.value,\n                     \"content\": content.value,\n                     \"image_base64\": articleImage.src\n                  }\n\n                  fetch('/api/article', {\n                     method: 'POST',\n                     headers: {\n                        'Content-Type': 'application/json'\n                     },\n                     body: JSON.stringify(request)\n                  }).then(response => {\n                     if (response.status !== 201) {\n                        response.text().then(function(text) {\n                           alert(`Не получилось сохранить статью: ${text}`)\n                        })\n                        return\n                     }\n\n                     location.reload();\n                  }).catch(err => alert(`Не получилось отправить запрос на сервер ${err}`))\n               }\n\n               metricsBtn.onclick = () => {\n               alert(\"works!\")\n               }\n            };\n          </script></head><body class=\"container\"><br><nav><ul><li><h1>Добро пожаловать в AI-feed!</h1></li></ul><ul><li><a href=\"/articles\">Статьи</a></li><li><a href=\"/themes\">Темы</a></li><li><a href=\"/personalities\">Личности</a></li></ul></nav><br><div><button class=\"container\" id=\"create-article\">Создать новую статью</button><br><br><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, article := range articles {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<article class=\"article\" data-uuid=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(article.ID.String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 377, Col: 74}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><details open><summary><div class=\"grid\"><h5><mark>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(article.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 381, Col: 49}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</mark></h5><div class=\"grid\"><button class=\"open\">Открыть</button> <button class=\"outline secondary delete\">Удалить</button> <button class=\"outline edit\">Редактировать</button></div></div></summary><h5><mark>Аналитика:</mark></h5><ul><li><h6>Количество слов - <i>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(article.WordsCount)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 391, Col: 84}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</i></h6></li><li><h6>Количество символов - <i>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(article.SymbolsCount)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 392, Col: 94}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</i></h6></li><li><div><h6>Ключевые слова</h6>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for index, el := range article.Keywords {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(index + 1))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 397, Col: 55}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(". ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(el.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 397, Col: 66}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" - <i>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(el.Count)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 397, Col: 82}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</i></p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></li></ul></details></article>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><dialog id=\"create-dialog\"><article><header><button aria-label=\"close\" rel=\"prev\" id=\"close-dialog\"></button><p><strong>Новая статья</strong></p></header><p><u>Создание новой статьи</u></p><form id=\"create-form\"><label>Тема <select id=\"select-theme\" name=\"theme\" aria-label=\"Выберите тему...\" required><option selected disabled value=\"\">Выберите тему...</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, theme := range themes {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(theme.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 426, Col: 51}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label> <label>Автор <select id=\"select-personality\" name=\"personality\" aria-label=\"Выберите автора...\" required><option selected disabled value=\"\">Выберите личность...</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, personality := range personalities {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option data-bio=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(personality.Biography)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 437, Col: 64}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-keywords=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(personality.Keywords)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 438, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-thematics=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(personality.Thematics)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 439, Col: 65}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-text-style=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(personality.TextStyle)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 440, Col: 66}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(personality.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/views/articles.templ`, Line: 441, Col: 46}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label> <label>Заголовок статьи <input id=\"article-title\" type=\"text\" name=\"title\" placeholder=\"Кликбейтный заголовок 1\" value=\"\"></label> <label>Текст статьи <textarea id=\"article-content\" name=\"content\" rows=\"10\" placeholder=\"Напишите или сгенерируйте статью\" aria-label=\"Напишите или сгенерируйте статью\"></textarea></label> <button id=\"generate-content\" class=\"outline\">Сгенерировать статью </button><br><br><img id=\"article-image\"><br><br><button id=\"generate-image\" class=\"outline\">Сгенерировать изображение</button><br></form><footer><button id=\"save-article-btn\">Создать</button></footer></article></dialog></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
