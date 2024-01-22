# \DefaultApi

All URIs are relative to *https://api.hh.ru*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicantComment**](DefaultApi.md#AddApplicantComment) | **Post** /applicant_comments/{applicant_id} | Добавление комментария
[**AddEmployerManager**](DefaultApi.md#AddEmployerManager) | **Post** /employers/{employer_id}/managers | Добавление менеджера
[**AddEmployerToBlacklisted**](DefaultApi.md#AddEmployerToBlacklisted) | **Put** /employers/blacklisted/{employer_id} | Добавление работодателя в список скрытых
[**AddResumeVisibilityList**](DefaultApi.md#AddResumeVisibilityList) | **Post** /resumes/{resume_id}/{list_type} | Добавление работодателей в список видимости
[**AddVacancyToArchive**](DefaultApi.md#AddVacancyToArchive) | **Put** /employers/{employer_id}/vacancies/archived/{vacancy_id} | Архивация вакансии
[**AddVacancyToBlacklisted**](DefaultApi.md#AddVacancyToBlacklisted) | **Put** /vacancies/blacklisted/{vacancy_id} | Добавление вакансии в список скрытых
[**AddVacancyToFavorite**](DefaultApi.md#AddVacancyToFavorite) | **Put** /vacancies/favorited/{vacancy_id} | Добавление вакансии в список отобранных
[**AddVacancyToHidden**](DefaultApi.md#AddVacancyToHidden) | **Put** /employers/{employer_id}/vacancies/hidden/{vacancy_id} | Удаление вакансий
[**ApplyToVacancy**](DefaultApi.md#ApplyToVacancy) | **Post** /negotiations | Отклик на вакансию
[**Authorize**](DefaultApi.md#Authorize) | **Post** /oauth/token | Получение access-токена
[**Authorize_0**](DefaultApi.md#Authorize_0) | **Post** /oauth/token | Получение access-токена
[**Authorize_1**](DefaultApi.md#Authorize_1) | **Post** /oauth/token | Получение access-токена
[**ChangeVacancyDraft**](DefaultApi.md#ChangeVacancyDraft) | **Put** /vacancies/drafts/{draft_id} | Изменение черновика вакансии
[**ConfirmPhoneInResume**](DefaultApi.md#ConfirmPhoneInResume) | **Post** /resume_phone_confirm | Подтвердить телефон кодом
[**CreateResume**](DefaultApi.md#CreateResume) | **Post** /resumes | Создание резюме
[**CreateSavedResumeSearch**](DefaultApi.md#CreateSavedResumeSearch) | **Post** /saved_searches/resumes | Создание нового сохраненного поиска резюме
[**CreateSavedVacancySearch**](DefaultApi.md#CreateSavedVacancySearch) | **Post** /saved_searches/vacancies | Создание нового сохраненного поиска вакансий
[**CreateVacancyDraft**](DefaultApi.md#CreateVacancyDraft) | **Post** /vacancies/drafts | Создание черновика вакансии
[**DeleteApplicantComment**](DefaultApi.md#DeleteApplicantComment) | **Delete** /applicant_comments/{applicant_id}/{comment_id} | Удаление комментария
[**DeleteArtifact**](DefaultApi.md#DeleteArtifact) | **Delete** /artifacts/{id} | Удаление артефакта
[**DeleteEmployerFromBlacklisted**](DefaultApi.md#DeleteEmployerFromBlacklisted) | **Delete** /employers/blacklisted/{employer_id} | Удаление работодателя из списка скрытых
[**DeleteEmployerFromResumeVisibilityList**](DefaultApi.md#DeleteEmployerFromResumeVisibilityList) | **Delete** /resumes/{resume_id}/{list_type}/employer | Удаление работодателя из списка видимости
[**DeleteEmployerManager**](DefaultApi.md#DeleteEmployerManager) | **Delete** /employers/{employer_id}/managers/{manager_id} | Удаление менеджера
[**DeleteResume**](DefaultApi.md#DeleteResume) | **Delete** /resumes/{resume_id} | Удаление резюме
[**DeleteResumeVisibilityList**](DefaultApi.md#DeleteResumeVisibilityList) | **Delete** /resumes/{resume_id}/{list_type} | Очистка списка видимости
[**DeleteSavedResumeSearch**](DefaultApi.md#DeleteSavedResumeSearch) | **Delete** /saved_searches/resumes/{id} | Удаление сохраненного поиска резюме
[**DeleteSavedVacancySearch**](DefaultApi.md#DeleteSavedVacancySearch) | **Delete** /saved_searches/vacancies/{id} | Удаление сохраненного поиска вакансий
[**DeleteVacancyDraft**](DefaultApi.md#DeleteVacancyDraft) | **Delete** /vacancies/drafts/{draft_id} | Удаление черновика вакансии
[**DeleteVacancyFromBlacklisted**](DefaultApi.md#DeleteVacancyFromBlacklisted) | **Delete** /vacancies/blacklisted/{vacancy_id} | Удаление вакансии из списка скрытых
[**DeleteVacancyFromFavorite**](DefaultApi.md#DeleteVacancyFromFavorite) | **Delete** /vacancies/favorited/{vacancy_id} | Удаление вакансии из списка отобранных
[**DisableAutomaticVacancyPublication**](DefaultApi.md#DisableAutomaticVacancyPublication) | **Delete** /vacancies/auto_publication | Отмена автопубликации вакансии
[**EditArtifact**](DefaultApi.md#EditArtifact) | **Put** /artifacts/{id} | Редактирование артефакта
[**EditCurrentUserInfo**](DefaultApi.md#EditCurrentUserInfo) | **Post** /me | Редактирование информации авторизованного пользователя
[**EditEmployerManager**](DefaultApi.md#EditEmployerManager) | **Put** /employers/{employer_id}/managers/{manager_id} | Редактирование менеджера
[**EditNegotiationMessage**](DefaultApi.md#EditNegotiationMessage) | **Put** /negotiations/{nid}/messages/{mid} | Редактирование сообщения в отклике
[**EditResume**](DefaultApi.md#EditResume) | **Put** /resumes/{resume_id} | Обновление резюме
[**EditVacancy**](DefaultApi.md#EditVacancy) | **Put** /vacancies/{vacancy_id} | Редактирование вакансий
[**GetActiveVacancyList**](DefaultApi.md#GetActiveVacancyList) | **Get** /employers/{employer_id}/vacancies/active | Просмотр списка опубликованных вакансий
[**GetActiveVacancyList_0**](DefaultApi.md#GetActiveVacancyList_0) | **Get** /employers/{employer_id}/vacancies/active | Просмотр списка опубликованных вакансий
[**GetAddress**](DefaultApi.md#GetAddress) | **Get** /employers/{employer_id}/addresses/{address_id} | Получение адреса
[**GetApplicantCommentsList**](DefaultApi.md#GetApplicantCommentsList) | **Get** /applicant_comments/{applicant_id} | Получение списка комментариев
[**GetApplicantPhoneInfo**](DefaultApi.md#GetApplicantPhoneInfo) | **Get** /resume_should_send_sms | Получить информацию о телефоне соискателя
[**GetArchivedVacancies**](DefaultApi.md#GetArchivedVacancies) | **Get** /employers/{employer_id}/vacancies/archived | Список архивных вакансий
[**GetArchivedVacancies_0**](DefaultApi.md#GetArchivedVacancies_0) | **Get** /employers/{employer_id}/vacancies/archived | Список архивных вакансий
[**GetAreaLeavesSuggests**](DefaultApi.md#GetAreaLeavesSuggests) | **Get** /suggests/area_leaves | Подсказки по регионам, являющимися листами в дереве регионов
[**GetAreas**](DefaultApi.md#GetAreas) | **Get** /areas | Дерево всех регионов
[**GetAreasFromSpecified**](DefaultApi.md#GetAreasFromSpecified) | **Get** /areas/{area_id} | Справочник регионов, начиная с указанного
[**GetAreasSuggests**](DefaultApi.md#GetAreasSuggests) | **Get** /suggests/areas | Подсказки по регионам
[**GetArtifactPhotos**](DefaultApi.md#GetArtifactPhotos) | **Get** /artifacts/photo | Получение фотографий
[**GetArtifactPhotosConditions**](DefaultApi.md#GetArtifactPhotosConditions) | **Get** /artifacts/photo/conditions | Условия загрузки фотографий
[**GetArtifactsPortfolio**](DefaultApi.md#GetArtifactsPortfolio) | **Get** /artifacts/portfolio | Получение портфолио
[**GetArtifactsPortfolioConditions**](DefaultApi.md#GetArtifactsPortfolioConditions) | **Get** /artifacts/portfolio/conditions | Условия загрузки портфолио
[**GetAvailableUserStatuses**](DefaultApi.md#GetAvailableUserStatuses) | **Get** /user_statuses/available | Справочник доступных пользовательских статусов
[**GetAvailableUserStatuses_0**](DefaultApi.md#GetAvailableUserStatuses_0) | **Get** /user_statuses/available | Справочник доступных пользовательских статусов
[**GetAvailableVacancyTypes**](DefaultApi.md#GetAvailableVacancyTypes) | **Get** /employers/{employer_id}/managers/{manager_id}/vacancies/available_types | Варианты публикации вакансий у текущего менеджера
[**GetBlacklistedEmployers**](DefaultApi.md#GetBlacklistedEmployers) | **Get** /employers/blacklisted | Список скрытых работодателей
[**GetBlacklistedVacancies**](DefaultApi.md#GetBlacklistedVacancies) | **Get** /vacancies/blacklisted | Список скрытых вакансий
[**GetCountries**](DefaultApi.md#GetCountries) | **Get** /areas/countries | Справочник стран
[**GetCurrentUserInfo**](DefaultApi.md#GetCurrentUserInfo) | **Get** /me | Информация о текущем пользователе
[**GetCurrentUserInfo_0**](DefaultApi.md#GetCurrentUserInfo_0) | **Get** /me | Информация о текущем пользователе
[**GetCurrentUserInfo_1**](DefaultApi.md#GetCurrentUserInfo_1) | **Get** /me | Информация о текущем пользователе
[**GetCurrentUserInfo_2**](DefaultApi.md#GetCurrentUserInfo_2) | **Get** /me | Информация о текущем пользователе
[**GetCurrentUserInfo_3**](DefaultApi.md#GetCurrentUserInfo_3) | **Get** /me | Информация о текущем пользователе
[**GetDictionaries**](DefaultApi.md#GetDictionaries) | **Get** /dictionaries | Справочники полей
[**GetEducationalInstitutionsDictionary**](DefaultApi.md#GetEducationalInstitutionsDictionary) | **Get** /educational_institutions | Основная информация об учебных заведениях
[**GetEducationalInstitutionsSuggests**](DefaultApi.md#GetEducationalInstitutionsSuggests) | **Get** /suggests/educational_institutions | Подсказки по названиям учебных заведений
[**GetEmployerAddresses**](DefaultApi.md#GetEmployerAddresses) | **Get** /employers/{employer_id}/addresses | Список адресов работодателя
[**GetEmployerDepartments**](DefaultApi.md#GetEmployerDepartments) | **Get** /employers/{employer_id}/departments | Справочник департаментов работодателя
[**GetEmployerDepartments_0**](DefaultApi.md#GetEmployerDepartments_0) | **Get** /employers/{employer_id}/departments | Справочник департаментов работодателя
[**GetEmployerInfo**](DefaultApi.md#GetEmployerInfo) | **Get** /employers/{employer_id} | Информация о работодателе
[**GetEmployerManager**](DefaultApi.md#GetEmployerManager) | **Get** /employers/{employer_id}/managers/{manager_id} | Получение информации о менеджере
[**GetEmployerManagerLimits**](DefaultApi.md#GetEmployerManagerLimits) | **Get** /employers/{employer_id}/managers/{manager_id}/limits/resume | Дневной лимит просмотра резюме для текущего менеджера
[**GetEmployerManagerTypes**](DefaultApi.md#GetEmployerManagerTypes) | **Get** /employers/{employer_id}/manager_types | Справочник типов и прав менеджера
[**GetEmployerManagerTypes_0**](DefaultApi.md#GetEmployerManagerTypes_0) | **Get** /employers/{employer_id}/manager_types | Справочник типов и прав менеджера
[**GetEmployerManager_0**](DefaultApi.md#GetEmployerManager_0) | **Get** /employers/{employer_id}/managers/{manager_id} | Получение информации о менеджере
[**GetEmployerManagers**](DefaultApi.md#GetEmployerManagers) | **Get** /employers/{employer_id}/managers | Список менеджеров работодателя
[**GetEmployerManagers_0**](DefaultApi.md#GetEmployerManagers_0) | **Get** /employers/{employer_id}/managers | Список менеджеров работодателя
[**GetEmployerVacancyAreas**](DefaultApi.md#GetEmployerVacancyAreas) | **Get** /employers/{employer_id}/vacancy_areas/active | Список регионов, в которых есть активные вакансии
[**GetFaculties**](DefaultApi.md#GetFaculties) | **Get** /educational_institutions/{id}/faculties | Список факультетов учебного заведения
[**GetFavoriteVacancies**](DefaultApi.md#GetFavoriteVacancies) | **Get** /vacancies/favorited | Список отобранных вакансий
[**GetFieldsOfStudySuggestions**](DefaultApi.md#GetFieldsOfStudySuggestions) | **Get** /suggests/fields_of_study | Подсказки по специализациям
[**GetHiddenVacancies**](DefaultApi.md#GetHiddenVacancies) | **Get** /employers/{employer_id}/vacancies/hidden | Список удаленных вакансий
[**GetHiddenVacancies_0**](DefaultApi.md#GetHiddenVacancies_0) | **Get** /employers/{employer_id}/vacancies/hidden | Список удаленных вакансий
[**GetIndustries**](DefaultApi.md#GetIndustries) | **Get** /industries | Отрасли компаний
[**GetLanguages**](DefaultApi.md#GetLanguages) | **Get** /languages | Список всех языков
[**GetLocales**](DefaultApi.md#GetLocales) | **Get** /locales | Список доступных локалей
[**GetLocalesForResume**](DefaultApi.md#GetLocalesForResume) | **Get** /locales/resume | Список доступных локалей для резюме
[**GetMailTemplates**](DefaultApi.md#GetMailTemplates) | **Get** /employers/{employer_id}/mail_templates | Список доступных шаблонов ответов соискателю
[**GetManagerAccounts**](DefaultApi.md#GetManagerAccounts) | **Get** /manager_accounts/mine | Рабочие аккаунты менеджера
[**GetManagerSettings**](DefaultApi.md#GetManagerSettings) | **Get** /employers/{employer_id}/managers/{manager_id}/settings | Предпочтения менеджера
[**GetMetroStations**](DefaultApi.md#GetMetroStations) | **Get** /metro | Список станций метро во всех городах
[**GetMetroStationsInCity**](DefaultApi.md#GetMetroStationsInCity) | **Get** /metro/{city_id} | Список станций метро в указанном городе
[**GetMineResumes**](DefaultApi.md#GetMineResumes) | **Get** /resumes/mine | Список резюме авторизованного пользователя
[**GetNewResumeConditions**](DefaultApi.md#GetNewResumeConditions) | **Get** /resume_conditions | Условия заполнения полей нового резюме
[**GetPayableApiActions**](DefaultApi.md#GetPayableApiActions) | **Get** /employers/{employer_id}/services/payable_api_actions/active | Информация по активным услугам API для платных методов
[**GetPositionsSuggestions**](DefaultApi.md#GetPositionsSuggestions) | **Get** /suggests/positions | Подсказки по должностям резюме
[**GetPrefNegotiationsOrder**](DefaultApi.md#GetPrefNegotiationsOrder) | **Get** /vacancies/{id}/preferred_negotiations_order | Просмотр предпочитаемой сортировки откликов
[**GetProfessionalRolesDictionary**](DefaultApi.md#GetProfessionalRolesDictionary) | **Get** /professional_roles | Справочник профессиональных ролей
[**GetProfessionalRolesSuggests**](DefaultApi.md#GetProfessionalRolesSuggests) | **Get** /suggests/professional_roles | Подсказки по профессиональным ролям
[**GetProlongationVacancyInfo**](DefaultApi.md#GetProlongationVacancyInfo) | **Get** /vacancies/{vacancy_id}/prolongate | Информация о возможности продления вакансии
[**GetRegisteredCompaniesSuggests**](DefaultApi.md#GetRegisteredCompaniesSuggests) | **Get** /suggests/companies | Подсказки по зарегистрированным организациям
[**GetResume**](DefaultApi.md#GetResume) | **Get** /resumes/{resume_id} | Просмотр резюме
[**GetResumeAccessTypes**](DefaultApi.md#GetResumeAccessTypes) | **Get** /resumes/{resume_id}/access_types | Получение списка типов видимости резюме
[**GetResumeConditions**](DefaultApi.md#GetResumeConditions) | **Get** /resumes/{resume_id}/conditions | Условия заполнения полей существующего резюме
[**GetResumeCreationAvailability**](DefaultApi.md#GetResumeCreationAvailability) | **Get** /resumes/creation_availability | Проверка возможности создания резюме
[**GetResumeNegotiationsHistory**](DefaultApi.md#GetResumeNegotiationsHistory) | **Get** /resumes/{resume_id}/negotiations_history | История откликов/приглашений по резюме
[**GetResumeSearchKeywordsSuggests**](DefaultApi.md#GetResumeSearchKeywordsSuggests) | **Get** /suggests/resume_search_keyword | Подсказки по ключевым словам поиска резюме
[**GetResumeStatus**](DefaultApi.md#GetResumeStatus) | **Get** /resumes/{resume_id}/status | Статус резюме и готовность к публикации
[**GetResumeViewHistory**](DefaultApi.md#GetResumeViewHistory) | **Get** /resumes/{resume_id}/views | История просмотра резюме
[**GetResumeVisibilityEmployersList**](DefaultApi.md#GetResumeVisibilityEmployersList) | **Get** /resumes/{resume_id}/{list_type}/search | Поиск работодателей для добавления в список видимости
[**GetResumeVisibilityList**](DefaultApi.md#GetResumeVisibilityList) | **Get** /resumes/{resume_id}/{list_type} | Получение списка видимости резюме
[**GetResume_0**](DefaultApi.md#GetResume_0) | **Get** /resumes/{resume_id} | Просмотр резюме
[**GetResume_1**](DefaultApi.md#GetResume_1) | **Get** /resumes/{resume_id} | Просмотр резюме
[**GetResumesByStatus**](DefaultApi.md#GetResumesByStatus) | **Get** /vacancies/{vacancy_id}/resumes_by_status | Резюме, сгруппированные по возможности отклика на данную вакансию
[**GetSalaryEmployeeLevels**](DefaultApi.md#GetSalaryEmployeeLevels) | **Get** /salary_statistics/dictionaries/employee_levels | Уровни компетенций
[**GetSalaryEvaluation**](DefaultApi.md#GetSalaryEvaluation) | **Get** /salary_statistics/paid/salary_evaluation/{area_id} | Оценка заработной платы без прогноза
[**GetSalaryEvaluation_0**](DefaultApi.md#GetSalaryEvaluation_0) | **Get** /salary_statistics/paid/salary_evaluation/{area_id} | Оценка заработной платы без прогноза
[**GetSalaryIndustries**](DefaultApi.md#GetSalaryIndustries) | **Get** /salary_statistics/dictionaries/salary_industries | Отрасли и сферы деятельности
[**GetSalaryProfessionalAreas**](DefaultApi.md#GetSalaryProfessionalAreas) | **Get** /salary_statistics/dictionaries/professional_areas | Профобласти и специализации
[**GetSalarySalaryAreas**](DefaultApi.md#GetSalarySalaryAreas) | **Get** /salary_statistics/dictionaries/salary_areas | Регионы и города
[**GetSavedResumeSearch**](DefaultApi.md#GetSavedResumeSearch) | **Get** /saved_searches/resumes/{id} | Получение единичного сохраненного поиска резюме
[**GetSavedResumeSearches**](DefaultApi.md#GetSavedResumeSearches) | **Get** /saved_searches/resumes | Список сохраненных поисков резюме
[**GetSavedVacancySearch**](DefaultApi.md#GetSavedVacancySearch) | **Get** /saved_searches/vacancies/{id} | Получение единичного сохраненного поиска вакансий
[**GetSavedVacancySearches**](DefaultApi.md#GetSavedVacancySearches) | **Get** /saved_searches/vacancies | Список сохраненных поисков вакансий
[**GetSkillSetSuggests**](DefaultApi.md#GetSkillSetSuggests) | **Get** /suggests/skill_set | Подсказки по ключевым навыкам
[**GetSkills**](DefaultApi.md#GetSkills) | **Get** /skills | Справочник ключевых навыков
[**GetSuitableResumes**](DefaultApi.md#GetSuitableResumes) | **Get** /vacancies/{vacancy_id}/suitable_resumes | Список подходящих для отклика резюме
[**GetTestsDictionary**](DefaultApi.md#GetTestsDictionary) | **Get** /employers/{employer_id}/tests | Справочник тестов работодателя
[**GetTestsDictionary_0**](DefaultApi.md#GetTestsDictionary_0) | **Get** /employers/{employer_id}/tests | Справочник тестов работодателя
[**GetVacancies**](DefaultApi.md#GetVacancies) | **Get** /vacancies | Поиск по вакансиям
[**GetVacanciesSimilarToResume**](DefaultApi.md#GetVacanciesSimilarToResume) | **Get** /resumes/{resume_id}/similar_vacancies | Поиск по вакансиям, похожим на резюме
[**GetVacanciesSimilarToVacancy**](DefaultApi.md#GetVacanciesSimilarToVacancy) | **Get** /vacancies/{vacancy_id}/similar_vacancies | Поиск по вакансиям, похожим на вакансию
[**GetVacancy**](DefaultApi.md#GetVacancy) | **Get** /vacancies/{vacancy_id} | Просмотр вакансии
[**GetVacancyBrandedTemplatesList**](DefaultApi.md#GetVacancyBrandedTemplatesList) | **Get** /employers/{employer_id}/vacancy_branded_templates | Список доступных бренд шаблонов вакансий работодателя
[**GetVacancyBrandedTemplatesList_0**](DefaultApi.md#GetVacancyBrandedTemplatesList_0) | **Get** /employers/{employer_id}/vacancy_branded_templates | Список доступных бренд шаблонов вакансий работодателя
[**GetVacancyConditions**](DefaultApi.md#GetVacancyConditions) | **Get** /vacancy_conditions | Условия заполнения полей при добавлении и редактировании вакансий
[**GetVacancyDraft**](DefaultApi.md#GetVacancyDraft) | **Get** /vacancies/drafts/{draft_id} | Получение черновика вакансии
[**GetVacancyDraftList**](DefaultApi.md#GetVacancyDraftList) | **Get** /vacancies/drafts | Получение списка черновиков вакансий
[**GetVacancyPositionsSuggests**](DefaultApi.md#GetVacancyPositionsSuggests) | **Get** /suggests/vacancy_positions | Подсказки по должностям вакансий
[**GetVacancySearchKeywords**](DefaultApi.md#GetVacancySearchKeywords) | **Get** /suggests/vacancy_search_keyword | Подсказки по ключевым словам поиска вакансий
[**GetVacancyStats**](DefaultApi.md#GetVacancyStats) | **Get** /vacancies/{vacancy_id}/stats | Статистика по вакансии
[**GetVacancyUpgradeList**](DefaultApi.md#GetVacancyUpgradeList) | **Get** /vacancies/{vacancy_id}/upgrades | Список улучшений для вакансии
[**GetVacancyVisitors**](DefaultApi.md#GetVacancyVisitors) | **Get** /vacancies/{vacancy_id}/visitors | Посмотревшие вакансию
[**GetVacancy_0**](DefaultApi.md#GetVacancy_0) | **Get** /vacancies/{vacancy_id} | Просмотр вакансии
[**HideActiveResponse**](DefaultApi.md#HideActiveResponse) | **Delete** /negotiations/active/{nid} | Скрыть отклик
[**InvalidateToken**](DefaultApi.md#InvalidateToken) | **Delete** /oauth/token | Инвалидация токена
[**InvalidateToken_0**](DefaultApi.md#InvalidateToken_0) | **Delete** /oauth/token | Инвалидация токена
[**InvalidateToken_1**](DefaultApi.md#InvalidateToken_1) | **Delete** /oauth/token | Инвалидация токена
[**LoadArtifact**](DefaultApi.md#LoadArtifact) | **Post** /artifacts | Загрузка артефакта
[**MoveSavedResumeSearch**](DefaultApi.md#MoveSavedResumeSearch) | **Put** /saved_searches/resumes/{saved_search_id}/managers/{manager_id} | Передача сохраненного поиска резюме другому менеджеру
[**PostNegotiationsTopicsRead**](DefaultApi.md#PostNegotiationsTopicsRead) | **Post** /negotiations/read | Отметить отклики прочитанными
[**PublishResume**](DefaultApi.md#PublishResume) | **Post** /resumes/{resume_id}/publish | Публикация резюме
[**PublishVacancy**](DefaultApi.md#PublishVacancy) | **Post** /vacancies | Публикация вакансии
[**PublishVacancyFromDraft**](DefaultApi.md#PublishVacancyFromDraft) | **Post** /vacancies/drafts/{draft_id}/publish | Публикация вакансии на основе черновика
[**PutMailTemplatesItem**](DefaultApi.md#PutMailTemplatesItem) | **Put** /employers/{employer_id}/mail_templates/{template_id} | Изменение шаблона ответа соискателю
[**PutNegotiationsCollectionToNextState**](DefaultApi.md#PutNegotiationsCollectionToNextState) | **Put** /negotiations/{collection} | Действия по откликам/приглашениям
[**PutPrefNegotiationsOrder**](DefaultApi.md#PutPrefNegotiationsOrder) | **Put** /vacancies/{id}/preferred_negotiations_order | Изменение предпочитаемой сортировки откликов
[**RestoreVacancyFromHidden**](DefaultApi.md#RestoreVacancyFromHidden) | **Delete** /employers/{employer_id}/vacancies/hidden/{vacancy_id} | Восстановление вакансии из удаленных
[**SearchEmployer**](DefaultApi.md#SearchEmployer) | **Get** /employers | Поиск работодателя
[**SearchForResumes**](DefaultApi.md#SearchForResumes) | **Get** /resumes | Поиск резюме
[**SearchForVacancyDraftDuplicates**](DefaultApi.md#SearchForVacancyDraftDuplicates) | **Get** /vacancies/drafts/{draft_id}/duplicates | Проверка наличия дубликатов вакансии
[**SendCodeForVerifyPhoneInResume**](DefaultApi.md#SendCodeForVerifyPhoneInResume) | **Post** /resume_phone_generate_code | Отправить код подтверждения для телефона резюме
[**SendNegotiationMessage**](DefaultApi.md#SendNegotiationMessage) | **Post** /negotiations/{nid}/messages | Отправка нового сообщения
[**UpdateApplicantComment**](DefaultApi.md#UpdateApplicantComment) | **Put** /applicant_comments/{applicant_id}/{comment_id} | Обновление комментария
[**UpdateSavedResumeSearch**](DefaultApi.md#UpdateSavedResumeSearch) | **Put** /saved_searches/resumes/{id} | Обновление сохраненного поиска резюме
[**UpdateSavedVacancySearch**](DefaultApi.md#UpdateSavedVacancySearch) | **Put** /saved_searches/vacancies/{id} | Обновление сохраненного поиска вакансий
[**VacancyProlongation**](DefaultApi.md#VacancyProlongation) | **Post** /vacancies/{vacancy_id}/prolongate | Продление вакансии



## AddApplicantComment

> ApplicantCommentsApplicantCommentItem AddApplicantComment(ctx, applicantId).HHUserAgent(hHUserAgent).Text(text).AccessType(accessType).Locale(locale).Host(host).Execute()

Добавление комментария



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	applicantId := "applicantId_example" // string | Идентификатор соискателя, который можно узнать из поля `owner` [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	text := "text_example" // string | Текст комментария
	accessType := "accessType_example" // string | Тип доступа. Доступные значения перечислены [в справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле `applicant_comment_access_type`
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.AddApplicantComment(context.Background(), applicantId).HHUserAgent(hHUserAgent).Text(text).AccessType(accessType).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddApplicantComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddApplicantComment`: ApplicantCommentsApplicantCommentItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddApplicantComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** | Идентификатор соискателя, который можно узнать из поля &#x60;owner&#x60; [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicantCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **text** | **string** | Текст комментария | 
 **accessType** | **string** | Тип доступа. Доступные значения перечислены [в справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле &#x60;applicant_comment_access_type&#x60; | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ApplicantCommentsApplicantCommentItem**](ApplicantCommentsApplicantCommentItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddEmployerManager

> EmployerManagersEmployerManagerId AddEmployerManager(ctx, employerId).HHUserAgent(hHUserAgent).EmployerManagersAddEmployerManager(employerManagersAddEmployerManager).Locale(locale).Host(host).Execute()

Добавление менеджера

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя. Чтобы получить его, используйте метод [Информация о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	employerManagersAddEmployerManager := *openapiclient.NewEmployerManagersAddEmployerManager(*openapiclient.NewEmployerManagersAreaId("Id_example"), "Email_example", "FirstName_example", false, "LastName_example", *openapiclient.NewEmployerManagersManagerTypeId("Id_example"), *openapiclient.NewEmployerManagersManagerDataPhone("City_example", "Country_example", "Number_example"), "Position_example") // EmployerManagersAddEmployerManager | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.AddEmployerManager(context.Background(), employerId).HHUserAgent(hHUserAgent).EmployerManagersAddEmployerManager(employerManagersAddEmployerManager).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddEmployerManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddEmployerManager`: EmployerManagersEmployerManagerId
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddEmployerManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя. Чтобы получить его, используйте метод [Информация о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEmployerManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **employerManagersAddEmployerManager** | [**EmployerManagersAddEmployerManager**](EmployerManagersAddEmployerManager.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagersEmployerManagerId**](EmployerManagersEmployerManagerId.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddEmployerToBlacklisted

> AddEmployerToBlacklisted(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Добавление работодателя в список скрытых



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.AddEmployerToBlacklisted(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddEmployerToBlacklisted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEmployerToBlacklistedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddResumeVisibilityList

> AddResumeVisibilityList(ctx, resumeId, listType).HHUserAgent(hHUserAgent).ResumesPostResumeVisibilityListBody(resumesPostResumeVisibilityListBody).Locale(locale).Host(host).Execute()

Добавление работодателей в список видимости



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	listType := "listType_example" // string | Тип списка. Допустимые значения — `whitelist` или `blacklist`
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	resumesPostResumeVisibilityListBody := *openapiclient.NewResumesPostResumeVisibilityListBody() // ResumesPostResumeVisibilityListBody | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.AddResumeVisibilityList(context.Background(), resumeId, listType).HHUserAgent(hHUserAgent).ResumesPostResumeVisibilityListBody(resumesPostResumeVisibilityListBody).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddResumeVisibilityList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 
**listType** | **string** | Тип списка. Допустимые значения — &#x60;whitelist&#x60; или &#x60;blacklist&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResumeVisibilityListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **resumesPostResumeVisibilityListBody** | [**ResumesPostResumeVisibilityListBody**](ResumesPostResumeVisibilityListBody.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVacancyToArchive

> AddVacancyToArchive(ctx, employerId, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Архивация вакансии



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.AddVacancyToArchive(context.Background(), employerId, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddVacancyToArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVacancyToArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVacancyToBlacklisted

> AddVacancyToBlacklisted(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Добавление вакансии в список скрытых



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.AddVacancyToBlacklisted(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddVacancyToBlacklisted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVacancyToBlacklistedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVacancyToFavorite

> AddVacancyToFavorite(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Добавление вакансии в список отобранных



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.AddVacancyToFavorite(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddVacancyToFavorite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVacancyToFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVacancyToHidden

> AddVacancyToHidden(ctx, employerId, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.AddVacancyToHidden(context.Background(), employerId, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddVacancyToHidden``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVacancyToHiddenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplyToVacancy

> string ApplyToVacancy(ctx).HHUserAgent(hHUserAgent).ResumeId(resumeId).VacancyId(vacancyId).Locale(locale).Host(host).Message(message).Execute()

Отклик на вакансию



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	resumeId := "resumeId_example" // string | Идентификатор резюме, которым производится отклик
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии, на которую происходит отклик
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")
	message := "message_example" // string | Сопроводительное письмо к отклику.  Является обязательным, если в вакансии указано, что обязательно сопроводительное письмо.  Максимальная длина — 10000 символов  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ApplyToVacancy(context.Background()).HHUserAgent(hHUserAgent).ResumeId(resumeId).VacancyId(vacancyId).Locale(locale).Host(host).Message(message).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApplyToVacancy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyToVacancy`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApplyToVacancy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyToVacancyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **resumeId** | **string** | Идентификатор резюме, которым производится отклик | 
 **vacancyId** | **string** | Идентификатор вакансии, на которую происходит отклик | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]
 **message** | **string** | Сопроводительное письмо к отклику.  Является обязательным, если в вакансии указано, что обязательно сопроводительное письмо.  Максимальная длина — 10000 символов  | 

### Return type

**string**

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Authorize

> AuthUserTokenAndAppToken Authorize(ctx).ClientId(clientId).ClientSecret(clientSecret).Code(code).GrantType(grantType).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()

Получение access-токена

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	clientId := "clientId_example" // string | Идентификатор, полученный при [создании приложения](https://dev.hh.ru/admin) (optional)
	clientSecret := "clientSecret_example" // string | Защищенный ключ, полученный при [создании приложения](https://dev.hh.ru/admin) (optional)
	code := "code_example" // string | Значение `authorization_code`, полученное при [перенаправлении пользователя](#get-authorization_code)  (optional)
	grantType := "grantType_example" // string | Cпособ запроса токена (optional)
	redirectUri := "redirectUri_example" // string | Uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра  (optional)
	refreshToken := "refreshToken_example" // string | Refresh-токен, полученный ранее при [получении пары токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) или прошлом обновлении пары (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.Authorize(context.Background()).ClientId(clientId).ClientSecret(clientSecret).Code(code).GrantType(grantType).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Authorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Authorize`: AuthUserTokenAndAppToken
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Authorize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Идентификатор, полученный при [создании приложения](https://dev.hh.ru/admin) | 
 **clientSecret** | **string** | Защищенный ключ, полученный при [создании приложения](https://dev.hh.ru/admin) | 
 **code** | **string** | Значение &#x60;authorization_code&#x60;, полученное при [перенаправлении пользователя](#get-authorization_code)  | 
 **grantType** | **string** | Cпособ запроса токена | 
 **redirectUri** | **string** | Uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра  | 
 **refreshToken** | **string** | Refresh-токен, полученный ранее при [получении пары токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) или прошлом обновлении пары | 

### Return type

[**AuthUserTokenAndAppToken**](AuthUserTokenAndAppToken.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Authorize_0

> AuthUserTokenAndAppToken Authorize_0(ctx).ClientId(clientId).ClientSecret(clientSecret).Code(code).GrantType(grantType).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()

Получение access-токена

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	clientId := "clientId_example" // string | Идентификатор, полученный при [создании приложения](https://dev.hh.ru/admin) (optional)
	clientSecret := "clientSecret_example" // string | Защищенный ключ, полученный при [создании приложения](https://dev.hh.ru/admin) (optional)
	code := "code_example" // string | Значение `authorization_code`, полученное при [перенаправлении пользователя](#get-authorization_code)  (optional)
	grantType := "grantType_example" // string | Cпособ запроса токена (optional)
	redirectUri := "redirectUri_example" // string | Uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра  (optional)
	refreshToken := "refreshToken_example" // string | Refresh-токен, полученный ранее при [получении пары токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) или прошлом обновлении пары (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.Authorize_0(context.Background()).ClientId(clientId).ClientSecret(clientSecret).Code(code).GrantType(grantType).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Authorize_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Authorize_0`: AuthUserTokenAndAppToken
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Authorize_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorize_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Идентификатор, полученный при [создании приложения](https://dev.hh.ru/admin) | 
 **clientSecret** | **string** | Защищенный ключ, полученный при [создании приложения](https://dev.hh.ru/admin) | 
 **code** | **string** | Значение &#x60;authorization_code&#x60;, полученное при [перенаправлении пользователя](#get-authorization_code)  | 
 **grantType** | **string** | Cпособ запроса токена | 
 **redirectUri** | **string** | Uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра  | 
 **refreshToken** | **string** | Refresh-токен, полученный ранее при [получении пары токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) или прошлом обновлении пары | 

### Return type

[**AuthUserTokenAndAppToken**](AuthUserTokenAndAppToken.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Authorize_1

> AuthUserTokenAndAppToken Authorize_1(ctx).ClientId(clientId).ClientSecret(clientSecret).Code(code).GrantType(grantType).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()

Получение access-токена

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	clientId := "clientId_example" // string | Идентификатор, полученный при [создании приложения](https://dev.hh.ru/admin) (optional)
	clientSecret := "clientSecret_example" // string | Защищенный ключ, полученный при [создании приложения](https://dev.hh.ru/admin) (optional)
	code := "code_example" // string | Значение `authorization_code`, полученное при [перенаправлении пользователя](#get-authorization_code)  (optional)
	grantType := "grantType_example" // string | Cпособ запроса токена (optional)
	redirectUri := "redirectUri_example" // string | Uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра  (optional)
	refreshToken := "refreshToken_example" // string | Refresh-токен, полученный ранее при [получении пары токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) или прошлом обновлении пары (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.Authorize_1(context.Background()).ClientId(clientId).ClientSecret(clientSecret).Code(code).GrantType(grantType).RedirectUri(redirectUri).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Authorize_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Authorize_1`: AuthUserTokenAndAppToken
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Authorize_1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorize_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Идентификатор, полученный при [создании приложения](https://dev.hh.ru/admin) | 
 **clientSecret** | **string** | Защищенный ключ, полученный при [создании приложения](https://dev.hh.ru/admin) | 
 **code** | **string** | Значение &#x60;authorization_code&#x60;, полученное при [перенаправлении пользователя](#get-authorization_code)  | 
 **grantType** | **string** | Cпособ запроса токена | 
 **redirectUri** | **string** | Uri для перенаправления пользователя после авторизации. Если не указать, используется из настроек приложения. При наличии происходит валидация значения. Вероятнее всего, потребуется сделать urlencode значения параметра  | 
 **refreshToken** | **string** | Refresh-токен, полученный ранее при [получении пары токенов](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov) или прошлом обновлении пары | 

### Return type

[**AuthUserTokenAndAppToken**](AuthUserTokenAndAppToken.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeVacancyDraft

> VacancyDraftDraftResponseSchema ChangeVacancyDraft(ctx, draftId).HHUserAgent(hHUserAgent).VacancyDraftVacancyDraftCreate(vacancyDraftVacancyDraftCreate).Locale(locale).Host(host).Execute()

Изменение черновика вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	draftId := "draftId_example" // string | Идентификатор черновика
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	vacancyDraftVacancyDraftCreate := *openapiclient.NewVacancyDraftVacancyDraftCreate() // VacancyDraftVacancyDraftCreate | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ChangeVacancyDraft(context.Background(), draftId).HHUserAgent(hHUserAgent).VacancyDraftVacancyDraftCreate(vacancyDraftVacancyDraftCreate).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ChangeVacancyDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeVacancyDraft`: VacancyDraftDraftResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ChangeVacancyDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string** | Идентификатор черновика | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeVacancyDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **vacancyDraftVacancyDraftCreate** | [**VacancyDraftVacancyDraftCreate**](VacancyDraftVacancyDraftCreate.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacancyDraftDraftResponseSchema**](VacancyDraftDraftResponseSchema.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmPhoneInResume

> ConfirmPhoneInResume(ctx).HHUserAgent(hHUserAgent).Phone(phone).ConfirmationCode(confirmationCode).Locale(locale).Host(host).Execute()

Подтвердить телефон кодом

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	phone := "phone_example" // string | Телефон который надо подтвердить
	confirmationCode := "confirmationCode_example" // string | Код для подтверждения
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.ConfirmPhoneInResume(context.Background()).HHUserAgent(hHUserAgent).Phone(phone).ConfirmationCode(confirmationCode).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConfirmPhoneInResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPhoneInResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **phone** | **string** | Телефон который надо подтвердить | 
 **confirmationCode** | **string** | Код для подтверждения | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResume

> string CreateResume(ctx).HHUserAgent(hHUserAgent).SourceResumeId(sourceResumeId).Locale(locale).Host(host).ResumeAddResumeRequest(resumeAddResumeRequest).Execute()

Создание резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	sourceResumeId := "sourceResumeId_example" // string | Идентификатор исходного резюме для клонирования (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")
	resumeAddResumeRequest := *openapiclient.NewResumeAddResumeRequest() // ResumeAddResumeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateResume(context.Background()).HHUserAgent(hHUserAgent).SourceResumeId(sourceResumeId).Locale(locale).Host(host).ResumeAddResumeRequest(resumeAddResumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResume`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateResume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **sourceResumeId** | **string** | Идентификатор исходного резюме для клонирования | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]
 **resumeAddResumeRequest** | [**ResumeAddResumeRequest**](ResumeAddResumeRequest.md) |  | 

### Return type

**string**

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSavedResumeSearch

> string CreateSavedResumeSearch(ctx).HHUserAgent(hHUserAgent).Text(text).TextLogic(textLogic).TextField(textField).TextPeriod(textPeriod).AgeFrom(ageFrom).AgeTo(ageTo).Area(area).Relocation(relocation).Period(period).DateFrom(dateFrom).DateTo(dateTo).EducationLevel(educationLevel).Employment(employment).Experience(experience).Skill(skill).Gender(gender).Label(label).Language(language).LanguageLevel(languageLevel).Metro(metro).Currency(currency).SalaryFrom(salaryFrom).SalaryTo(salaryTo).Schedule(schedule).OrderBy(orderBy).Citizenship(citizenship).WorkTicket(workTicket).EducationalInstitution(educationalInstitution).SearchInResponses(searchInResponses).ByTextPrefix(byTextPrefix).DriverLicenseTypes(driverLicenseTypes).VacancyId(vacancyId).Page(page).PerPage(perPage).ProfessionalRole(professionalRole).Locale(locale).Host(host).Execute()

Создание нового сохраненного поиска резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	text := "text_example" // string | Поисковая фраза. Метод найдет резюме, в которых встречаются все слова заданной фразы.  Особенности:  * Можно указать несколько значений. Каждое дополнительное значение уточняет поиск. * В качестве поисковой фразы можно использовать [язык поисковых запросов](http://hh.ru/article.xml?articleId=1175). * Специально для этого поля предусмотрено [автодополнение по подсказкам](#tag/Podskazki/operation/get-resume-search-keywords-suggests). * Для тонкой настройки по фразе можно использовать параметры `text.logic`, `text.field`, `text.period`. При использовании дополнительных `text.*` полей, необходимо указывать весь набор (триаду) параметров  (optional)
	textLogic := "textLogic_example" // string | Описывает, как производится поиск. Возможные значения перечислены в поле `resume_search_logic` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) (optional)
	textField := "textField_example" // string | Описывает, где должны встречаться слова из поисковой фразы `text`. Можно указать несколько значений через запятую, например `?text.field=education,keywords`. Возможные значения перечислены в поле `resume_search_fields` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) (optional)
	textPeriod := "textPeriod_example" // string | Период опыта работы.  Параметр имеет смысл только при `text.field` равным одному из значений: `experience`, `experience_company`, `experience_position`, `experience_description`, но указывать его необходимо всегда при указании других `text.*`. Если параметр не имеет смысла, то его значение можно оставить пустым  (optional)
	ageFrom := float32(8.14) // float32 | Нижняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение `only_with_age` в параметре `label`  (optional)
	ageTo := float32(8.14) // float32 | Верхняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение `only_with_age` в параметре `label`  (optional)
	area := "area_example" // string | Регион. Возможные значения указаны в [справочнике регионов](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений.  По умолчанию выбираются резюме, в которых соискатели живут в указанных регионах или готовы в них переехать. Поменять это поведение поиска можно, указав параметр `relocation`  (optional)
	relocation := "relocation_example" // string | Готовность к переезду. Возможные значения указаны в поле `resume_search_relocation` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Необходимо указывать вместе с параметром `area`  (optional)
	period := float32(8.14) // float32 | Поиск ведется по резюме, опубликованным за указанный период в днях. Если период не указан, поиск ведется без ограничений по дате публикации  (optional)
	dateFrom := "dateFrom_example" // string | Дата, от которой нужно начать поиск. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — `YYYY-MM-DD` или с точностью до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Нельзя передавать вместе с параметром `period`  (optional)
	dateTo := "dateTo_example" // string | Дата, до которой нужно искать. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — `YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Можно передавать только в паре с параметром `date_from`. Нельзя передавать вместе с параметром `period`  (optional)
	educationLevel := "educationLevel_example" // string | Уровень образования. Возможные значения перечислены в поле `education_level` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Если параметр не указан, поиск ведется без ограничений на уровень образования  (optional)
	employment := "employment_example" // string | Тип занятости. Возможные значения перечислены в поле `employment` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	experience := "experience_example" // string | Опыт работы. Возможные значения перечислены в поле `experience` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	skill := "skill_example" // string | Ключевые навыки. Указывается один или несколько идентификаторов ключевых навыков. Значения можно получить из поля `id` в [подсказке по ключевым навыкам](#tag/Podskazki/operation/get-skill-set-suggests)  (optional)
	gender := "gender_example" // string | Пол соискателя. Возможные значения перечислены в поле `gender` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries).  По умолчанию вне зависимости от значения параметра будут найдены резюме, в которых пол не указан, убрать такие резюме можно с помощью параметра `label=only_with_gender`  (optional)
	label := "label_example" // string | Дополнительный фильтр. Возможные значения перечислены в поле `resume_search_label` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	language := "language_example" // string | Знание языка. Можно указать несколько значений.  Возможные значения перечислены в [справочнике языков](#tag/Obshie-spravochniki/operation/get-languages)  (optional)
	languageLevel := "languageLevel_example" // string | Уровень знания языка. Можно указать несколько значений.  Возможные значения перечислены в поле `language_level` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	metro := "metro_example" // string | Линия, либо станция метро. Возможные значения перечислены в [справочнике метро](#tag/Obshie-spravochniki/operation/get-metro-stations)  (optional)
	currency := "currency_example" // string | Код валюты. Возможные значения перечислены в поле `currency.code` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	salaryFrom := float32(8.14) // float32 | Нижняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр `label=only_with_salary`  (optional)
	salaryTo := float32(8.14) // float32 | Верхняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр `label=only_with_salary`  (optional)
	schedule := "schedule_example" // string | График работы. Возможные значения перечислены в поле `schedule` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка резюме. Возможные значения перечислены в поле `resume_search_order` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	citizenship := "citizenship_example" // string | Страна гражданства соискателя. Возможные значения перечислены в [справочнике стран](#tag/Obshie-spravochniki/operation/get-countries). Можно указать несколько значений  (optional)
	workTicket := "workTicket_example" // string | Страна, в которой у соискателя есть разрешение на работу. Возможные значения перечислены в [справочнике стран](#tag/Obshie-spravochniki/operation/get-countries). Можно указать несколько значений  (optional)
	educationalInstitution := "educationalInstitution_example" // string | Учебные заведения соискателя. В качестве параметров используются [подсказки по названиям университетов](#tag/Podskazki). Можно указать несколько значений  (optional)
	searchInResponses := true // bool | Если `true`, то поиск осуществляется только по резюме, которыми соискатели откликались на вакансии компании текущего пользователя, если `false` — поиск осуществляется по всем резюме. По умолчанию — `false`  (optional)
	byTextPrefix := true // bool | Если `true`, включается поиск по префиксу. Для каждого параметра `text` будут находиться не только полные совпадения слов, но еще и слова, начинающиеся с `text`. По умолчанию — `false`  (optional)
	driverLicenseTypes := "driverLicenseTypes_example" // string | Категории водительских прав соискателя. Возможные значения перечислены в поле `driver_license_types` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии для поиска похожих резюме. Необходимо передавать идентификатор активной вакансии работодателя или вакансии работодателя в архиве  (optional)
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию — 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию — 10, максимальное значение — 50) (optional)
	professionalRole := "professionalRole_example" // string | Профессиональная роль. Элемент справочника [профессиональных ролей](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary). Можно указать несколько значений  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateSavedResumeSearch(context.Background()).HHUserAgent(hHUserAgent).Text(text).TextLogic(textLogic).TextField(textField).TextPeriod(textPeriod).AgeFrom(ageFrom).AgeTo(ageTo).Area(area).Relocation(relocation).Period(period).DateFrom(dateFrom).DateTo(dateTo).EducationLevel(educationLevel).Employment(employment).Experience(experience).Skill(skill).Gender(gender).Label(label).Language(language).LanguageLevel(languageLevel).Metro(metro).Currency(currency).SalaryFrom(salaryFrom).SalaryTo(salaryTo).Schedule(schedule).OrderBy(orderBy).Citizenship(citizenship).WorkTicket(workTicket).EducationalInstitution(educationalInstitution).SearchInResponses(searchInResponses).ByTextPrefix(byTextPrefix).DriverLicenseTypes(driverLicenseTypes).VacancyId(vacancyId).Page(page).PerPage(perPage).ProfessionalRole(professionalRole).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSavedResumeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSavedResumeSearch`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSavedResumeSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSavedResumeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **text** | **string** | Поисковая фраза. Метод найдет резюме, в которых встречаются все слова заданной фразы.  Особенности:  * Можно указать несколько значений. Каждое дополнительное значение уточняет поиск. * В качестве поисковой фразы можно использовать [язык поисковых запросов](http://hh.ru/article.xml?articleId&#x3D;1175). * Специально для этого поля предусмотрено [автодополнение по подсказкам](#tag/Podskazki/operation/get-resume-search-keywords-suggests). * Для тонкой настройки по фразе можно использовать параметры &#x60;text.logic&#x60;, &#x60;text.field&#x60;, &#x60;text.period&#x60;. При использовании дополнительных &#x60;text.*&#x60; полей, необходимо указывать весь набор (триаду) параметров  | 
 **textLogic** | **string** | Описывает, как производится поиск. Возможные значения перечислены в поле &#x60;resume_search_logic&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
 **textField** | **string** | Описывает, где должны встречаться слова из поисковой фразы &#x60;text&#x60;. Можно указать несколько значений через запятую, например &#x60;?text.field&#x3D;education,keywords&#x60;. Возможные значения перечислены в поле &#x60;resume_search_fields&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
 **textPeriod** | **string** | Период опыта работы.  Параметр имеет смысл только при &#x60;text.field&#x60; равным одному из значений: &#x60;experience&#x60;, &#x60;experience_company&#x60;, &#x60;experience_position&#x60;, &#x60;experience_description&#x60;, но указывать его необходимо всегда при указании других &#x60;text.*&#x60;. Если параметр не имеет смысла, то его значение можно оставить пустым  | 
 **ageFrom** | **float32** | Нижняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение &#x60;only_with_age&#x60; в параметре &#x60;label&#x60;  | 
 **ageTo** | **float32** | Верхняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение &#x60;only_with_age&#x60; в параметре &#x60;label&#x60;  | 
 **area** | **string** | Регион. Возможные значения указаны в [справочнике регионов](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений.  По умолчанию выбираются резюме, в которых соискатели живут в указанных регионах или готовы в них переехать. Поменять это поведение поиска можно, указав параметр &#x60;relocation&#x60;  | 
 **relocation** | **string** | Готовность к переезду. Возможные значения указаны в поле &#x60;resume_search_relocation&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Необходимо указывать вместе с параметром &#x60;area&#x60;  | 
 **period** | **float32** | Поиск ведется по резюме, опубликованным за указанный период в днях. Если период не указан, поиск ведется без ограничений по дате публикации  | 
 **dateFrom** | **string** | Дата, от которой нужно начать поиск. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — &#x60;YYYY-MM-DD&#x60; или с точностью до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Нельзя передавать вместе с параметром &#x60;period&#x60;  | 
 **dateTo** | **string** | Дата, до которой нужно искать. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — &#x60;YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Можно передавать только в паре с параметром &#x60;date_from&#x60;. Нельзя передавать вместе с параметром &#x60;period&#x60;  | 
 **educationLevel** | **string** | Уровень образования. Возможные значения перечислены в поле &#x60;education_level&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Если параметр не указан, поиск ведется без ограничений на уровень образования  | 
 **employment** | **string** | Тип занятости. Возможные значения перечислены в поле &#x60;employment&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **experience** | **string** | Опыт работы. Возможные значения перечислены в поле &#x60;experience&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **skill** | **string** | Ключевые навыки. Указывается один или несколько идентификаторов ключевых навыков. Значения можно получить из поля &#x60;id&#x60; в [подсказке по ключевым навыкам](#tag/Podskazki/operation/get-skill-set-suggests)  | 
 **gender** | **string** | Пол соискателя. Возможные значения перечислены в поле &#x60;gender&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries).  По умолчанию вне зависимости от значения параметра будут найдены резюме, в которых пол не указан, убрать такие резюме можно с помощью параметра &#x60;label&#x3D;only_with_gender&#x60;  | 
 **label** | **string** | Дополнительный фильтр. Возможные значения перечислены в поле &#x60;resume_search_label&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **language** | **string** | Знание языка. Можно указать несколько значений.  Возможные значения перечислены в [справочнике языков](#tag/Obshie-spravochniki/operation/get-languages)  | 
 **languageLevel** | **string** | Уровень знания языка. Можно указать несколько значений.  Возможные значения перечислены в поле &#x60;language_level&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **metro** | **string** | Линия, либо станция метро. Возможные значения перечислены в [справочнике метро](#tag/Obshie-spravochniki/operation/get-metro-stations)  | 
 **currency** | **string** | Код валюты. Возможные значения перечислены в поле &#x60;currency.code&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **salaryFrom** | **float32** | Нижняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр &#x60;label&#x3D;only_with_salary&#x60;  | 
 **salaryTo** | **float32** | Верхняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр &#x60;label&#x3D;only_with_salary&#x60;  | 
 **schedule** | **string** | График работы. Возможные значения перечислены в поле &#x60;schedule&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **orderBy** | **string** | Сортировка списка резюме. Возможные значения перечислены в поле &#x60;resume_search_order&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **citizenship** | **string** | Страна гражданства соискателя. Возможные значения перечислены в [справочнике стран](#tag/Obshie-spravochniki/operation/get-countries). Можно указать несколько значений  | 
 **workTicket** | **string** | Страна, в которой у соискателя есть разрешение на работу. Возможные значения перечислены в [справочнике стран](#tag/Obshie-spravochniki/operation/get-countries). Можно указать несколько значений  | 
 **educationalInstitution** | **string** | Учебные заведения соискателя. В качестве параметров используются [подсказки по названиям университетов](#tag/Podskazki). Можно указать несколько значений  | 
 **searchInResponses** | **bool** | Если &#x60;true&#x60;, то поиск осуществляется только по резюме, которыми соискатели откликались на вакансии компании текущего пользователя, если &#x60;false&#x60; — поиск осуществляется по всем резюме. По умолчанию — &#x60;false&#x60;  | 
 **byTextPrefix** | **bool** | Если &#x60;true&#x60;, включается поиск по префиксу. Для каждого параметра &#x60;text&#x60; будут находиться не только полные совпадения слов, но еще и слова, начинающиеся с &#x60;text&#x60;. По умолчанию — &#x60;false&#x60;  | 
 **driverLicenseTypes** | **string** | Категории водительских прав соискателя. Возможные значения перечислены в поле &#x60;driver_license_types&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **vacancyId** | **string** | Идентификатор вакансии для поиска похожих резюме. Необходимо передавать идентификатор активной вакансии работодателя или вакансии работодателя в архиве  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию — 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию — 10, максимальное значение — 50) | 
 **professionalRole** | **string** | Профессиональная роль. Элемент справочника [профессиональных ролей](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary). Можно указать несколько значений  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

**string**

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSavedVacancySearch

> string CreateSavedVacancySearch(ctx).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).Locale(locale).Host(host).Execute()

Создание нового сохраненного поиска вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию - 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию - 10, максимальное значение - 100) (optional)
	text := "text_example" // string | Переданное значение ищется в полях вакансии, указанных в параметре `search_field`. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) (optional)
	searchField := "searchField_example" // string | Область поиска. Справочник с возможными значениями: `vacancy_search_fields` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  (optional)
	experience := "experience_example" // string | Опыт работы. Необходимо передавать `id` из справочника `experience` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	employment := "employment_example" // string | Тип занятости. Необходимо передавать `id` из справочника `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	schedule := "schedule_example" // string | График работы. Необходимо передавать `id` из справочника `schedule` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	area := "area_example" // string | Регион. Необходимо передавать `id` из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  (optional)
	metro := "metro_example" // string | Ветка или станция метро. Необходимо передавать `id` из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  (optional)
	professionalRole := "professionalRole_example" // string | Профессиональная область. Необходимо передавать `id` из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  (optional)
	industry := "industry_example" // string | Индустрия компании, разместившей вакансию. Необходимо передавать `id` из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  (optional)
	employerId := "employerId_example" // string | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  (optional)
	currency := "currency_example" // string | Код валюты. Справочник с возможными значениями: `currency` (ключ `code`) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром `salary`  (optional)
	salary := float32(8.14) // float32 | Размер заработной платы. Если указано это поле, но не указано `currency`, то для `currency` используется значение RUR.  При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании `salary=100&currency=EUR` будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR.  По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте `only_with_salary=true`  (optional)
	label := "label_example" // string | Фильтр по меткам вакансий. Необходимо передавать `id` из справочника `vacancy_label` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	onlyWithSalary := true // bool | Показывать вакансии только с указанием зарплаты. По умолчанию `false`  (optional)
	period := float32(8.14) // float32 | Количество дней, в пределах которых производится поиск по вакансиям  (optional)
	dateFrom := "dateFrom_example" // string | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`.  Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	dateTo := "dateTo_example" // string | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`.  Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	topLat := float32(8.14) // float32 | Верхняя граница широты.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	bottomLat := float32(8.14) // float32 | Нижняя граница широты.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	leftLng := float32(8.14) // float32 | Левая граница долготы.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	rightLng := float32(8.14) // float32 | Правая граница долготы.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий. Справочник с возможными значениями: `vacancy_search_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries).  Если выбрана сортировка по удалённости от гео-точки `distance`, необходимо также задать её координаты: `sort_point_lat`, `sort_point_lng`  (optional)
	sortPointLat := float32(8.14) // float32 | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	sortPointLng := float32(8.14) // float32 | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	clusters := true // bool | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — `false`  (optional)
	describeArguments := true // bool | Возвращать ли описание использованных параметров поиска. Успешный ответ будет содержать поле [`arguments`]((#tag/Poisk-vakansij/operation/get-vacancies))).  По умолчанию — `false`  (optional)
	noMagic := true // bool | Если значение `true` — автоматическое преобразование вакансий отключено. По умолчанию – false.  При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос `text=москва бухгалтер 100500` будет преобразован в `text=бухгалтер&only_with_salary=true&area=1&salary=100500`  (optional)
	premium := true // bool | Если значение `true` — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  (optional)
	responsesCountEnabled := true // bool | Если значение `true` — дополнительное поле `counters` с количеством откликов для вакансии включено. По-умолчанию — `false`  (optional)
	partTime := true // bool | Вакансии для подработки. Возможные значения:  * Все элементы из `working_days` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_intervals` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_modes` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы `part` или `project` из `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент `accept_temporary`, показывает вакансии только с временным трудоустройством.  Можно указать несколько значений  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateSavedVacancySearch(context.Background()).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSavedVacancySearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSavedVacancySearch`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSavedVacancySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSavedVacancySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию - 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию - 10, максимальное значение - 100) | 
 **text** | **string** | Переданное значение ищется в полях вакансии, указанных в параметре &#x60;search_field&#x60;. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) | 
 **searchField** | **string** | Область поиска. Справочник с возможными значениями: &#x60;vacancy_search_fields&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  | 
 **experience** | **string** | Опыт работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;experience&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **employment** | **string** | Тип занятости. Необходимо передавать &#x60;id&#x60; из справочника &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **schedule** | **string** | График работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;schedule&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **area** | **string** | Регион. Необходимо передавать &#x60;id&#x60; из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  | 
 **metro** | **string** | Ветка или станция метро. Необходимо передавать &#x60;id&#x60; из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  | 
 **professionalRole** | **string** | Профессиональная область. Необходимо передавать &#x60;id&#x60; из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  | 
 **industry** | **string** | Индустрия компании, разместившей вакансию. Необходимо передавать &#x60;id&#x60; из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  | 
 **employerId** | **string** | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  | 
 **currency** | **string** | Код валюты. Справочник с возможными значениями: &#x60;currency&#x60; (ключ &#x60;code&#x60;) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром &#x60;salary&#x60;  | 
 **salary** | **float32** | Размер заработной платы. Если указано это поле, но не указано &#x60;currency&#x60;, то для &#x60;currency&#x60; используется значение RUR.  При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании &#x60;salary&#x3D;100&amp;currency&#x3D;EUR&#x60; будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR.  По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте &#x60;only_with_salary&#x3D;true&#x60;  | 
 **label** | **string** | Фильтр по меткам вакансий. Необходимо передавать &#x60;id&#x60; из справочника &#x60;vacancy_label&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **onlyWithSalary** | **bool** | Показывать вакансии только с указанием зарплаты. По умолчанию &#x60;false&#x60;  | 
 **period** | **float32** | Количество дней, в пределах которых производится поиск по вакансиям  | 
 **dateFrom** | **string** | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;.  Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **dateTo** | **string** | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;.  Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **topLat** | **float32** | Верхняя граница широты.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **bottomLat** | **float32** | Нижняя граница широты.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **leftLng** | **float32** | Левая граница долготы.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **rightLng** | **float32** | Правая граница долготы.  При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби.  Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **orderBy** | **string** | Сортировка списка вакансий. Справочник с возможными значениями: &#x60;vacancy_search_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries).  Если выбрана сортировка по удалённости от гео-точки &#x60;distance&#x60;, необходимо также задать её координаты: &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;  | 
 **sortPointLat** | **float32** | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **sortPointLng** | **float32** | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **clusters** | **bool** | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — &#x60;false&#x60;  | 
 **describeArguments** | **bool** | Возвращать ли описание использованных параметров поиска. Успешный ответ будет содержать поле [&#x60;arguments&#x60;]((#tag/Poisk-vakansij/operation/get-vacancies))).  По умолчанию — &#x60;false&#x60;  | 
 **noMagic** | **bool** | Если значение &#x60;true&#x60; — автоматическое преобразование вакансий отключено. По умолчанию – false.  При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос &#x60;text&#x3D;москва бухгалтер 100500&#x60; будет преобразован в &#x60;text&#x3D;бухгалтер&amp;only_with_salary&#x3D;true&amp;area&#x3D;1&amp;salary&#x3D;100500&#x60;  | 
 **premium** | **bool** | Если значение &#x60;true&#x60; — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  | 
 **responsesCountEnabled** | **bool** | Если значение &#x60;true&#x60; — дополнительное поле &#x60;counters&#x60; с количеством откликов для вакансии включено. По-умолчанию — &#x60;false&#x60;  | 
 **partTime** | **bool** | Вакансии для подработки. Возможные значения:  * Все элементы из &#x60;working_days&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_intervals&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_modes&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы &#x60;part&#x60; или &#x60;project&#x60; из &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент &#x60;accept_temporary&#x60;, показывает вакансии только с временным трудоустройством.  Можно указать несколько значений  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

**string**

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVacancyDraft

> VacancyDraftDraftResponseSchema CreateVacancyDraft(ctx).HHUserAgent(hHUserAgent).VacancyDraftVacancyDraftCreate(vacancyDraftVacancyDraftCreate).Locale(locale).Host(host).Execute()

Создание черновика вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	vacancyDraftVacancyDraftCreate := *openapiclient.NewVacancyDraftVacancyDraftCreate() // VacancyDraftVacancyDraftCreate | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.CreateVacancyDraft(context.Background()).HHUserAgent(hHUserAgent).VacancyDraftVacancyDraftCreate(vacancyDraftVacancyDraftCreate).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVacancyDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVacancyDraft`: VacancyDraftDraftResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVacancyDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVacancyDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **vacancyDraftVacancyDraftCreate** | [**VacancyDraftVacancyDraftCreate**](VacancyDraftVacancyDraftCreate.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacancyDraftDraftResponseSchema**](VacancyDraftDraftResponseSchema.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicantComment

> DeleteApplicantComment(ctx, applicantId, commentId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление комментария



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	applicantId := "applicantId_example" // string | Идентификатор соискателя, который можно узнать из поля `owner` [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field)
	commentId := "commentId_example" // string | Идентификатор комментария, который будет удален. Его можно узнать в [списке комментариев](#tag/Kommentarii-k-soiskatelyu/operation/get-applicant-comments-list)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteApplicantComment(context.Background(), applicantId, commentId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplicantComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** | Идентификатор соискателя, который можно узнать из поля &#x60;owner&#x60; [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field) | 
**commentId** | **string** | Идентификатор комментария, который будет удален. Его можно узнать в [списке комментариев](#tag/Kommentarii-k-soiskatelyu/operation/get-applicant-comments-list) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicantCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifact

> DeleteArtifact(ctx, id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление артефакта

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор артефакта. Чтобы получить его, используйте метод [Получение портфолио](#tag/Rabota-s-artefaktami/operation/get-artifacts-portfolio) или [Получение фотографий](#tag/Rabota-s-artefaktami/operation/get-artifact-photos)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteArtifact(context.Background(), id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор артефакта. Чтобы получить его, используйте метод [Получение портфолио](#tag/Rabota-s-artefaktami/operation/get-artifacts-portfolio) или [Получение фотографий](#tag/Rabota-s-artefaktami/operation/get-artifact-photos) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmployerFromBlacklisted

> DeleteEmployerFromBlacklisted(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление работодателя из списка скрытых



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteEmployerFromBlacklisted(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEmployerFromBlacklisted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmployerFromBlacklistedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmployerFromResumeVisibilityList

> DeleteEmployerFromResumeVisibilityList(ctx, resumeId, listType).Id(id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление работодателя из списка видимости



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	listType := "listType_example" // string | Тип списка. Допустимые значения — `whitelist` или `blacklist`
	id := "id_example" // string | Идентификатор работодателя. Множественный параметр
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteEmployerFromResumeVisibilityList(context.Background(), resumeId, listType).Id(id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEmployerFromResumeVisibilityList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 
**listType** | **string** | Тип списка. Допустимые значения — &#x60;whitelist&#x60; или &#x60;blacklist&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmployerFromResumeVisibilityListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **string** | Идентификатор работодателя. Множественный параметр | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmployerManager

> DeleteEmployerManager(ctx, employerId, managerId).SuccessorId(successorId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление менеджера



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	managerId := "managerId_example" // string | Идентификатор менеджера
	successorId := "successorId_example" // string | Идентификатор менеджера, которому передаются данные, связанные с удаляемым менеджером, в частности: вакансии, отклики, папки отобранных резюме, комментарии к соискателю, автопоиски и прочее
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteEmployerManager(context.Background(), employerId, managerId).SuccessorId(successorId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEmployerManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**managerId** | **string** | Идентификатор менеджера | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmployerManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **successorId** | **string** | Идентификатор менеджера, которому передаются данные, связанные с удаляемым менеджером, в частности: вакансии, отклики, папки отобранных резюме, комментарии к соискателю, автопоиски и прочее | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResume

> DeleteResume(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteResume(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResumeVisibilityList

> DeleteResumeVisibilityList(ctx, resumeId, listType).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Очистка списка видимости

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	listType := "listType_example" // string | Тип списка. Допустимые значения — `whitelist` или `blacklist`
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteResumeVisibilityList(context.Background(), resumeId, listType).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteResumeVisibilityList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 
**listType** | **string** | Тип списка. Допустимые значения — &#x60;whitelist&#x60; или &#x60;blacklist&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResumeVisibilityListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSavedResumeSearch

> DeleteSavedResumeSearch(ctx, id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление сохраненного поиска резюме

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор сохраненного поиска
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteSavedResumeSearch(context.Background(), id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSavedResumeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор сохраненного поиска | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavedResumeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSavedVacancySearch

> DeleteSavedVacancySearch(ctx, id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление сохраненного поиска вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор сохраненного поиска
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteSavedVacancySearch(context.Background(), id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSavedVacancySearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор сохраненного поиска | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavedVacancySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVacancyDraft

> DeleteVacancyDraft(ctx, draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление черновика вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	draftId := "draftId_example" // string | Идентификатор черновика
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteVacancyDraft(context.Background(), draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVacancyDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string** | Идентификатор черновика | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVacancyDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVacancyFromBlacklisted

> DeleteVacancyFromBlacklisted(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление вакансии из списка скрытых



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteVacancyFromBlacklisted(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVacancyFromBlacklisted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVacancyFromBlacklistedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVacancyFromFavorite

> DeleteVacancyFromFavorite(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Удаление вакансии из списка отобранных



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteVacancyFromFavorite(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVacancyFromFavorite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVacancyFromFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAutomaticVacancyPublication

> DisableAutomaticVacancyPublication(ctx).DraftId(draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Отмена автопубликации вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	draftId := "draftId_example" // string | Идентификатор черновика
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DisableAutomaticVacancyPublication(context.Background()).DraftId(draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DisableAutomaticVacancyPublication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableAutomaticVacancyPublicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **draftId** | **string** | Идентификатор черновика | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditArtifact

> EditArtifact(ctx, id).HHUserAgent(hHUserAgent).Description(description).Locale(locale).Host(host).Execute()

Редактирование артефакта



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор артефакта. Чтобы получить его, используйте метод [Получение портфолио](#tag/Rabota-s-artefaktami/operation/get-artifacts-portfolio)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	description := "description_example" // string | Описание изображения
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.EditArtifact(context.Background(), id).HHUserAgent(hHUserAgent).Description(description).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор артефакта. Чтобы получить его, используйте метод [Получение портфолио](#tag/Rabota-s-artefaktami/operation/get-artifacts-portfolio) | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **description** | **string** | Описание изображения | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCurrentUserInfo

> EditCurrentUserInfo(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).FirstName(firstName).LastName(lastName).MiddleName(middleName).IsInSearch(isInSearch).Execute()

Редактирование информации авторизованного пользователя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")
	firstName := "firstName_example" // string | Имя (optional)
	lastName := "lastName_example" // string | Фамилия (optional)
	middleName := "middleName_example" // string | Отчество, поле может быть пустым (optional)
	isInSearch := "isInSearch_example" // string | Флаг «ищу работу». Возможные значения: `true`/`false` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.EditCurrentUserInfo(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).FirstName(firstName).LastName(lastName).MiddleName(middleName).IsInSearch(isInSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditCurrentUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditCurrentUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]
 **firstName** | **string** | Имя | 
 **lastName** | **string** | Фамилия | 
 **middleName** | **string** | Отчество, поле может быть пустым | 
 **isInSearch** | **string** | Флаг «ищу работу». Возможные значения: &#x60;true&#x60;/&#x60;false&#x60; | 

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditEmployerManager

> EditEmployerManager(ctx, employerId, managerId).HHUserAgent(hHUserAgent).EmployerManagersManagerData(employerManagersManagerData).Locale(locale).Host(host).Execute()

Редактирование менеджера

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	managerId := "managerId_example" // string | Идентификатор менеджера. Можно узнать из списка [менеджеров](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	employerManagersManagerData := map[string][]openapiclient.EmployerManagersManagerData{ ... } // EmployerManagersManagerData | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.EditEmployerManager(context.Background(), employerId, managerId).HHUserAgent(hHUserAgent).EmployerManagersManagerData(employerManagersManagerData).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditEmployerManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**managerId** | **string** | Идентификатор менеджера. Можно узнать из списка [менеджеров](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers) | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEmployerManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **employerManagersManagerData** | [**EmployerManagersManagerData**](EmployerManagersManagerData.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditNegotiationMessage

> EditNegotiationMessage(ctx, nid, mid).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Редактирование сообщения в отклике



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	nid := "nid_example" // string | Идентификатор отклика
	mid := "mid_example" // string | Идентификатор сообщения в отклике
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.EditNegotiationMessage(context.Background(), nid, mid).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditNegotiationMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nid** | **string** | Идентификатор отклика | 
**mid** | **string** | Идентификатор сообщения в отклике | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditNegotiationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditResume

> EditResume(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).ResumeEditResumeRequest(resumeEditResumeRequest).Execute()

Обновление резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")
	resumeEditResumeRequest := *openapiclient.NewResumeEditResumeRequest() // ResumeEditResumeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.EditResume(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).ResumeEditResumeRequest(resumeEditResumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]
 **resumeEditResumeRequest** | [**ResumeEditResumeRequest**](ResumeEditResumeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditVacancy

> EditVacancy(ctx, vacancyId).HHUserAgent(hHUserAgent).VacancyEdit(vacancyEdit).IgnoreDuplicates(ignoreDuplicates).Locale(locale).Host(host).Execute()

Редактирование вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	vacancyEdit := map[string][]openapiclient.VacancyEdit{ ... } // VacancyEdit | 
	ignoreDuplicates := true // bool | Игнорировать [появление дубликата](https://github.com/hhru/api/blob/master/docs/employer_vacancies.md#edit-ignore-duplicates), после редактирования вакансии. По умолчанию — `false` (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.EditVacancy(context.Background(), vacancyId).HHUserAgent(hHUserAgent).VacancyEdit(vacancyEdit).IgnoreDuplicates(ignoreDuplicates).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditVacancy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditVacancyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **vacancyEdit** | [**VacancyEdit**](VacancyEdit.md) |  | 
 **ignoreDuplicates** | **bool** | Игнорировать [появление дубликата](https://github.com/hhru/api/blob/master/docs/employer_vacancies.md#edit-ignore-duplicates), после редактирования вакансии. По умолчанию — &#x60;false&#x60; | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveVacancyList

> VacanciesVacancyListResponse GetActiveVacancyList(ctx, employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).ManagerId(managerId).Text(text).Area(area).ResumeId(resumeId).OrderBy(orderBy).Locale(locale).Host(host).Execute()

Просмотр списка опубликованных вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0) (optional) (default to 0)
	perPage := float32(8.14) // float32 | Количество элементов (optional) (default to 20)
	managerId := "managerId_example" // string | Идентификатор менеджера, вакансии которого будут получены в ответе. По умолчанию возвращаются вакансии текущего пользователя.  Если передать несколько `manager_id`, будет использован последний. Значения можно взять из [списка](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)  (optional)
	text := "text_example" // string | Строка для поиска по названию вакансии (optional)
	area := "area_example" // string | Идентификатор региона с вакансией. Чтобы получить идентификаторы регионов, в которых есть активные вакансии, воспользуйтесь [соответствующим методом](#tag/Informaciya-o-rabotodatele/operation/get-employer-vacancy-areas) (optional)
	resumeId := "resumeId_example" // string | Идентификатор резюме. Этот параметр нельзя передавать в комбинации с другими параметрами, только отдельно. Если параметр передан, в ответе возвращаются только те вакансии, которые подходят для указанного резюме, а также дополнительные поля (optional)
	orderBy := "orderBy_example" // string | Способ сортировки вакансий. Доступные значения перечислены в поле `employer_active_vacancies_order` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetActiveVacancyList(context.Background(), employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).ManagerId(managerId).Text(text).Area(area).ResumeId(resumeId).OrderBy(orderBy).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetActiveVacancyList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveVacancyList`: VacanciesVacancyListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetActiveVacancyList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveVacancyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0) | [default to 0]
 **perPage** | **float32** | Количество элементов | [default to 20]
 **managerId** | **string** | Идентификатор менеджера, вакансии которого будут получены в ответе. По умолчанию возвращаются вакансии текущего пользователя.  Если передать несколько &#x60;manager_id&#x60;, будет использован последний. Значения можно взять из [списка](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)  | 
 **text** | **string** | Строка для поиска по названию вакансии | 
 **area** | **string** | Идентификатор региона с вакансией. Чтобы получить идентификаторы регионов, в которых есть активные вакансии, воспользуйтесь [соответствующим методом](#tag/Informaciya-o-rabotodatele/operation/get-employer-vacancy-areas) | 
 **resumeId** | **string** | Идентификатор резюме. Этот параметр нельзя передавать в комбинации с другими параметрами, только отдельно. Если параметр передан, в ответе возвращаются только те вакансии, которые подходят для указанного резюме, а также дополнительные поля | 
 **orderBy** | **string** | Способ сортировки вакансий. Доступные значения перечислены в поле &#x60;employer_active_vacancies_order&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancyListResponse**](VacanciesVacancyListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveVacancyList_0

> VacanciesVacancyListResponse GetActiveVacancyList_0(ctx, employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).ManagerId(managerId).Text(text).Area(area).ResumeId(resumeId).OrderBy(orderBy).Locale(locale).Host(host).Execute()

Просмотр списка опубликованных вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0) (optional) (default to 0)
	perPage := float32(8.14) // float32 | Количество элементов (optional) (default to 20)
	managerId := "managerId_example" // string | Идентификатор менеджера, вакансии которого будут получены в ответе. По умолчанию возвращаются вакансии текущего пользователя.  Если передать несколько `manager_id`, будет использован последний. Значения можно взять из [списка](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)  (optional)
	text := "text_example" // string | Строка для поиска по названию вакансии (optional)
	area := "area_example" // string | Идентификатор региона с вакансией. Чтобы получить идентификаторы регионов, в которых есть активные вакансии, воспользуйтесь [соответствующим методом](#tag/Informaciya-o-rabotodatele/operation/get-employer-vacancy-areas) (optional)
	resumeId := "resumeId_example" // string | Идентификатор резюме. Этот параметр нельзя передавать в комбинации с другими параметрами, только отдельно. Если параметр передан, в ответе возвращаются только те вакансии, которые подходят для указанного резюме, а также дополнительные поля (optional)
	orderBy := "orderBy_example" // string | Способ сортировки вакансий. Доступные значения перечислены в поле `employer_active_vacancies_order` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetActiveVacancyList_0(context.Background(), employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).ManagerId(managerId).Text(text).Area(area).ResumeId(resumeId).OrderBy(orderBy).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetActiveVacancyList_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveVacancyList_0`: VacanciesVacancyListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetActiveVacancyList_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveVacancyList_3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0) | [default to 0]
 **perPage** | **float32** | Количество элементов | [default to 20]
 **managerId** | **string** | Идентификатор менеджера, вакансии которого будут получены в ответе. По умолчанию возвращаются вакансии текущего пользователя.  Если передать несколько &#x60;manager_id&#x60;, будет использован последний. Значения можно взять из [списка](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)  | 
 **text** | **string** | Строка для поиска по названию вакансии | 
 **area** | **string** | Идентификатор региона с вакансией. Чтобы получить идентификаторы регионов, в которых есть активные вакансии, воспользуйтесь [соответствующим методом](#tag/Informaciya-o-rabotodatele/operation/get-employer-vacancy-areas) | 
 **resumeId** | **string** | Идентификатор резюме. Этот параметр нельзя передавать в комбинации с другими параметрами, только отдельно. Если параметр передан, в ответе возвращаются только те вакансии, которые подходят для указанного резюме, а также дополнительные поля | 
 **orderBy** | **string** | Способ сортировки вакансий. Доступные значения перечислены в поле &#x60;employer_active_vacancies_order&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancyListResponse**](VacanciesVacancyListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddress

> EmployerAddressesEmployerAddressItemResponse GetAddress(ctx, employerId, addressId).HHUserAgent(hHUserAgent).WithManager(withManager).Locale(locale).Host(host).Execute()

Получение адреса

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя. Чтобы получить его, используйте метод [Информация о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	addressId := "addressId_example" // string | Идентификатор адреса работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	withManager := true // bool | Если true, ответ будет содержать информацию о менеджере создавшем адрес (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAddress(context.Background(), employerId, addressId).HHUserAgent(hHUserAgent).WithManager(withManager).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddress`: EmployerAddressesEmployerAddressItemResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя. Чтобы получить его, используйте метод [Информация о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**addressId** | **string** | Идентификатор адреса работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **withManager** | **bool** | Если true, ответ будет содержать информацию о менеджере создавшем адрес | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerAddressesEmployerAddressItemResponse**](EmployerAddressesEmployerAddressItemResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicantCommentsList

> ApplicantCommentsApplicantCommentsList GetApplicantCommentsList(ctx, applicantId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).OrderBy(orderBy).Locale(locale).Host(host).Execute()

Получение списка комментариев



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	applicantId := "applicantId_example" // string | Идентификатор соискателя, который можно узнать из поля `owner` [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (optional)
	perPage := float32(8.14) // float32 | Результатов на странице (optional)
	orderBy := "orderBy_example" // string | Сортировка комментариев. Доступные значения перечислены [в справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле `applicant_comments_order` (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetApplicantCommentsList(context.Background(), applicantId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).OrderBy(orderBy).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicantCommentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicantCommentsList`: ApplicantCommentsApplicantCommentsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicantCommentsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** | Идентификатор соискателя, который можно узнать из поля &#x60;owner&#x60; [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicantCommentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы | 
 **perPage** | **float32** | Результатов на странице | 
 **orderBy** | **string** | Сортировка комментариев. Доступные значения перечислены [в справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле &#x60;applicant_comments_order&#x60; | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ApplicantCommentsApplicantCommentsList**](ApplicantCommentsApplicantCommentsList.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicantPhoneInfo

> ResumeShouldSendSmsContainer GetApplicantPhoneInfo(ctx).Phone(phone).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получить информацию о телефоне соискателя

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	phone := "phone_example" // string | Номер телефона в любом формате
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetApplicantPhoneInfo(context.Background()).Phone(phone).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicantPhoneInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicantPhoneInfo`: ResumeShouldSendSmsContainer
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicantPhoneInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicantPhoneInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Номер телефона в любом формате | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumeShouldSendSmsContainer**](ResumeShouldSendSmsContainer.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivedVacancies

> VacanciesArchivedVacancyListResponse GetArchivedVacancies(ctx, employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()

Список архивных вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	managerId := "managerId_example" // string | Идентификатор менеджера из [списка менеджеров работодателя](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers). Передайте, если требуется получить вакансии другого менеджера.   Если передать несколько параметров `manager_id`, будет использоваться только последний.  По умолчанию возвращаются вакансии текущего пользователя  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий в архиве. Справочник с возможными значениями: `employer_archived_vacancies_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	perPage := int32(56) // int32 | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение `per_page` составляет 1000  (optional)
	page := int32(56) // int32 | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetArchivedVacancies(context.Background(), employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetArchivedVacancies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArchivedVacancies`: VacanciesArchivedVacancyListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetArchivedVacancies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedVacanciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **managerId** | **string** | Идентификатор менеджера из [списка менеджеров работодателя](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers). Передайте, если требуется получить вакансии другого менеджера.   Если передать несколько параметров &#x60;manager_id&#x60;, будет использоваться только последний.  По умолчанию возвращаются вакансии текущего пользователя  | 
 **orderBy** | **string** | Сортировка списка вакансий в архиве. Справочник с возможными значениями: &#x60;employer_archived_vacancies_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **perPage** | **int32** | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение &#x60;per_page&#x60; составляет 1000  | 
 **page** | **int32** | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesArchivedVacancyListResponse**](VacanciesArchivedVacancyListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivedVacancies_0

> VacanciesArchivedVacancyListResponse GetArchivedVacancies_0(ctx, employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()

Список архивных вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	managerId := "managerId_example" // string | Идентификатор менеджера из [списка менеджеров работодателя](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers). Передайте, если требуется получить вакансии другого менеджера.   Если передать несколько параметров `manager_id`, будет использоваться только последний.  По умолчанию возвращаются вакансии текущего пользователя  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий в архиве. Справочник с возможными значениями: `employer_archived_vacancies_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	perPage := int32(56) // int32 | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение `per_page` составляет 1000  (optional)
	page := int32(56) // int32 | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetArchivedVacancies_0(context.Background(), employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetArchivedVacancies_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArchivedVacancies_0`: VacanciesArchivedVacancyListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetArchivedVacancies_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedVacancies_4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **managerId** | **string** | Идентификатор менеджера из [списка менеджеров работодателя](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers). Передайте, если требуется получить вакансии другого менеджера.   Если передать несколько параметров &#x60;manager_id&#x60;, будет использоваться только последний.  По умолчанию возвращаются вакансии текущего пользователя  | 
 **orderBy** | **string** | Сортировка списка вакансий в архиве. Справочник с возможными значениями: &#x60;employer_archived_vacancies_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **perPage** | **int32** | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение &#x60;per_page&#x60; составляет 1000  | 
 **page** | **int32** | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesArchivedVacancyListResponse**](VacanciesArchivedVacancyListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAreaLeavesSuggests

> SuggestsAreas GetAreaLeavesSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).AreaId(areaId).Locale(locale).Host(host).Execute()

Подсказки по регионам, являющимися листами в дереве регионов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска региона. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	areaId := int32(56) // int32 | Идентификатор региона из [справочника](#tag/Obshie-spravochniki/operation/get-areas). Сужает подсказки поддеревом переданного идентификатора региона (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAreaLeavesSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).AreaId(areaId).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAreaLeavesSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAreaLeavesSuggests`: SuggestsAreas
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAreaLeavesSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAreaLeavesSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска региона. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **areaId** | **int32** | Идентификатор региона из [справочника](#tag/Obshie-spravochniki/operation/get-areas). Сужает подсказки поддеревом переданного идентификатора региона | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsAreas**](SuggestsAreas.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAreas

> []DictionariesAreaItem GetAreas(ctx).HHUserAgent(hHUserAgent).AdditionalCase(additionalCase).Locale(locale).Host(host).Execute()

Дерево всех регионов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	additionalCase := "additionalCase_example" // string | Применимо только к русской локализации.  В дополнительном поле вернется название региона в указанном падеже. Поддерживается только значение `prepositional` — предложный падеж  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAreas(context.Background()).HHUserAgent(hHUserAgent).AdditionalCase(additionalCase).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAreas`: []DictionariesAreaItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAreas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **additionalCase** | **string** | Применимо только к русской локализации.  В дополнительном поле вернется название региона в указанном падеже. Поддерживается только значение &#x60;prepositional&#x60; — предложный падеж  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]DictionariesAreaItem**](DictionariesAreaItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAreasFromSpecified

> DictionariesAreaItem GetAreasFromSpecified(ctx, areaId).HHUserAgent(hHUserAgent).AdditionalCase(additionalCase).Locale(locale).Host(host).Execute()

Справочник регионов, начиная с указанного



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	areaId := "areaId_example" // string | Идентификатор региона из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	additionalCase := "additionalCase_example" // string | Применимо только к русской локализации.  В дополнительном поле вернется название региона в указанном падеже. Поддерживается только значение `prepositional` — предложный падеж  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAreasFromSpecified(context.Background(), areaId).HHUserAgent(hHUserAgent).AdditionalCase(additionalCase).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAreasFromSpecified``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAreasFromSpecified`: DictionariesAreaItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAreasFromSpecified`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**areaId** | **string** | Идентификатор региона из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAreasFromSpecifiedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **additionalCase** | **string** | Применимо только к русской локализации.  В дополнительном поле вернется название региона в указанном падеже. Поддерживается только значение &#x60;prepositional&#x60; — предложный падеж  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**DictionariesAreaItem**](DictionariesAreaItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAreasSuggests

> SuggestsAreas GetAreasSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).AreaId(areaId).IncludeParent(includeParent).Locale(locale).Host(host).Execute()

Подсказки по регионам



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска региона. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	areaId := int32(56) // int32 | Идентификатор региона из [справочника](#tag/Obshie-spravochniki/operation/get-areas). Сужает подсказки поддеревом переданного идентификатора региона (optional)
	includeParent := true // bool | Включать ли в ответ регион, переданный в параметре `area_id`, если он подходит по искомому тексту (optional) (default to false)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAreasSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).AreaId(areaId).IncludeParent(includeParent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAreasSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAreasSuggests`: SuggestsAreas
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAreasSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAreasSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска региона. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **areaId** | **int32** | Идентификатор региона из [справочника](#tag/Obshie-spravochniki/operation/get-areas). Сужает подсказки поддеревом переданного идентификатора региона | 
 **includeParent** | **bool** | Включать ли в ответ регион, переданный в параметре &#x60;area_id&#x60;, если он подходит по искомому тексту | [default to false]
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsAreas**](SuggestsAreas.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactPhotos

> ArtifactsArtifactPhotoResponse GetArtifactPhotos(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение фотографий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetArtifactPhotos(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetArtifactPhotos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactPhotos`: ArtifactsArtifactPhotoResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetArtifactPhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactPhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ArtifactsArtifactPhotoResponse**](ArtifactsArtifactPhotoResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactPhotosConditions

> ArtifactsArtifactConditions GetArtifactPhotosConditions(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Условия загрузки фотографий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetArtifactPhotosConditions(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetArtifactPhotosConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactPhotosConditions`: ArtifactsArtifactConditions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetArtifactPhotosConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactPhotosConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ArtifactsArtifactConditions**](ArtifactsArtifactConditions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactsPortfolio

> ArtifactsArtifactPortfolioResponse GetArtifactsPortfolio(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение портфолио



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetArtifactsPortfolio(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetArtifactsPortfolio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactsPortfolio`: ArtifactsArtifactPortfolioResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetArtifactsPortfolio`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactsPortfolioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ArtifactsArtifactPortfolioResponse**](ArtifactsArtifactPortfolioResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactsPortfolioConditions

> ArtifactsArtifactConditions GetArtifactsPortfolioConditions(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Условия загрузки портфолио

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetArtifactsPortfolioConditions(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetArtifactsPortfolioConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactsPortfolioConditions`: ArtifactsArtifactConditions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetArtifactsPortfolioConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactsPortfolioConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ArtifactsArtifactConditions**](ArtifactsArtifactConditions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableUserStatuses

> UserStatusesAvailable GetAvailableUserStatuses(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник доступных пользовательских статусов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAvailableUserStatuses(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAvailableUserStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableUserStatuses`: UserStatusesAvailable
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAvailableUserStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableUserStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**UserStatusesAvailable**](UserStatusesAvailable.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableUserStatuses_0

> UserStatusesAvailable GetAvailableUserStatuses_0(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник доступных пользовательских статусов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAvailableUserStatuses_0(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAvailableUserStatuses_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableUserStatuses_0`: UserStatusesAvailable
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAvailableUserStatuses_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableUserStatuses_5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**UserStatusesAvailable**](UserStatusesAvailable.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableVacancyTypes

> VacanciesAvailableVacancyTypeResponse GetAvailableVacancyTypes(ctx, employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Варианты публикации вакансий у текущего менеджера



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	managerId := "managerId_example" // string | Идентификатор менеджера, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetAvailableVacancyTypes(context.Background(), employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAvailableVacancyTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableVacancyTypes`: VacanciesAvailableVacancyTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAvailableVacancyTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**managerId** | **string** | Идентификатор менеджера, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableVacancyTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesAvailableVacancyTypeResponse**](VacanciesAvailableVacancyTypeResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlacklistedEmployers

> EmployersEmployersBlacklistedResponse GetBlacklistedEmployers(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список скрытых работодателей



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetBlacklistedEmployers(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBlacklistedEmployers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlacklistedEmployers`: EmployersEmployersBlacklistedResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBlacklistedEmployers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlacklistedEmployersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersEmployersBlacklistedResponse**](EmployersEmployersBlacklistedResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlacklistedVacancies

> VacanciesVacanciesBlacklistedResponse GetBlacklistedVacancies(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список скрытых вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetBlacklistedVacancies(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBlacklistedVacancies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlacklistedVacancies`: VacanciesVacanciesBlacklistedResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBlacklistedVacancies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlacklistedVacanciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacanciesBlacklistedResponse**](VacanciesVacanciesBlacklistedResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> []IncludesArea GetCountries(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник стран



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetCountries(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: []IncludesArea
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]IncludesArea**](IncludesArea.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserInfo

> MeProfile GetCurrentUserInfo(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация о текущем пользователе

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetCurrentUserInfo(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserInfo`: MeProfile
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentUserInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**MeProfile**](MeProfile.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserInfo_0

> MeProfile GetCurrentUserInfo_0(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация о текущем пользователе

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetCurrentUserInfo_0(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentUserInfo_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserInfo_0`: MeProfile
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentUserInfo_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserInfo_6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**MeProfile**](MeProfile.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserInfo_1

> MeProfile GetCurrentUserInfo_1(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация о текущем пользователе

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetCurrentUserInfo_1(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentUserInfo_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserInfo_1`: MeProfile
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentUserInfo_1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserInfo_7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**MeProfile**](MeProfile.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserInfo_2

> MeProfile GetCurrentUserInfo_2(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация о текущем пользователе

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetCurrentUserInfo_2(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentUserInfo_2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserInfo_2`: MeProfile
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentUserInfo_2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserInfo_8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**MeProfile**](MeProfile.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserInfo_3

> MeProfile GetCurrentUserInfo_3(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация о текущем пользователе

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetCurrentUserInfo_3(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentUserInfo_3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserInfo_3`: MeProfile
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentUserInfo_3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserInfo_9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**MeProfile**](MeProfile.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDictionaries

> DictionariesDictResponse GetDictionaries(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочники полей



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetDictionaries(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDictionaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDictionaries`: DictionariesDictResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDictionaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDictionariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**DictionariesDictResponse**](DictionariesDictResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEducationalInstitutionsDictionary

> SuggestsEducationalInstitutions GetEducationalInstitutionsDictionary(ctx).Id(id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Основная информация об учебных заведениях

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := int32(56) // int32 | Идентификаторы учебных заведений. Идентификатор конкретного заведения можно узнать в [подсказке](#tag/Podskazki/operation/get-educational-institutions-suggests). Передать можно не более 50 значений. Например: `?id=39196&id=45470&id=0`. Если был передан идентификатор несуществующего заведения, для него не вернется никакой информации
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEducationalInstitutionsDictionary(context.Background()).Id(id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEducationalInstitutionsDictionary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEducationalInstitutionsDictionary`: SuggestsEducationalInstitutions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEducationalInstitutionsDictionary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEducationalInstitutionsDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | Идентификаторы учебных заведений. Идентификатор конкретного заведения можно узнать в [подсказке](#tag/Podskazki/operation/get-educational-institutions-suggests). Передать можно не более 50 значений. Например: &#x60;?id&#x3D;39196&amp;id&#x3D;45470&amp;id&#x3D;0&#x60;. Если был передан идентификатор несуществующего заведения, для него не вернется никакой информации | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsEducationalInstitutions**](SuggestsEducationalInstitutions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEducationalInstitutionsSuggests

> SuggestsEducationalInstitutions GetEducationalInstitutionsSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по названиям учебных заведений



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска учебного заведения. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEducationalInstitutionsSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEducationalInstitutionsSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEducationalInstitutionsSuggests`: SuggestsEducationalInstitutions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEducationalInstitutionsSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEducationalInstitutionsSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска учебного заведения. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsEducationalInstitutions**](SuggestsEducationalInstitutions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerAddresses

> EmployerAddressesEmployerAddressesResponse GetEmployerAddresses(ctx, employerId).HHUserAgent(hHUserAgent).ChangedAfter(changedAfter).ManagerId(managerId).WithManager(withManager).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()

Список адресов работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	changedAfter := "changedAfter_example" // string | Позволяет загрузить все адреса, изменённые после этой даты (добавление, удаление или изменение адреса). Изменения возвращаются без пагинации. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) - `YYYY-MM-DDThh:mm:ss` или c указанием отступа для часового пояса `YYYY-MM-DDThh:mm:ss±hhmm`. Максимальное значение отступа от текущей даты - 7 дней. При передаче этого параметра, для каждого адреса в теле ответа возвращается поле `deleted`, указывающее на то, удалён ли адрес (optional)
	managerId := "managerId_example" // string | Идентификатор менеджера создавшего адрес (optional)
	withManager := true // bool | Если true, ответ будет содержать информацию о менеджере создавшем адрес (optional)
	perPage := int32(56) // int32 | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение per_page составляет 10000  (optional)
	page := int32(56) // int32 | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerAddresses(context.Background(), employerId).HHUserAgent(hHUserAgent).ChangedAfter(changedAfter).ManagerId(managerId).WithManager(withManager).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerAddresses`: EmployerAddressesEmployerAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **changedAfter** | **string** | Позволяет загрузить все адреса, изменённые после этой даты (добавление, удаление или изменение адреса). Изменения возвращаются без пагинации. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) - &#x60;YYYY-MM-DDThh:mm:ss&#x60; или c указанием отступа для часового пояса &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Максимальное значение отступа от текущей даты - 7 дней. При передаче этого параметра, для каждого адреса в теле ответа возвращается поле &#x60;deleted&#x60;, указывающее на то, удалён ли адрес | 
 **managerId** | **string** | Идентификатор менеджера создавшего адрес | 
 **withManager** | **bool** | Если true, ответ будет содержать информацию о менеджере создавшем адрес | 
 **perPage** | **int32** | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение per_page составляет 10000  | 
 **page** | **int32** | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerAddressesEmployerAddressesResponse**](EmployerAddressesEmployerAddressesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerDepartments

> EmployersEmployerDepartmentsResponse GetEmployerDepartments(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник департаментов работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно получить в [списке работодателей](#tag/Rabotodatel/operation/search-employer)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerDepartments(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerDepartments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerDepartments`: EmployersEmployerDepartmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerDepartments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно получить в [списке работодателей](#tag/Rabotodatel/operation/search-employer) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerDepartmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersEmployerDepartmentsResponse**](EmployersEmployerDepartmentsResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerDepartments_0

> EmployersEmployerDepartmentsResponse GetEmployerDepartments_0(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник департаментов работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно получить в [списке работодателей](#tag/Rabotodatel/operation/search-employer)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerDepartments_0(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerDepartments_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerDepartments_0`: EmployersEmployerDepartmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerDepartments_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно получить в [списке работодателей](#tag/Rabotodatel/operation/search-employer) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerDepartments_10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersEmployerDepartmentsResponse**](EmployersEmployerDepartmentsResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerInfo

> EmployersEmployerInfo GetEmployerInfo(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация о работодателе



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно получить в [списке работодателей](#tag/Rabotodatel/operation/search-employer)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerInfo(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerInfo`: EmployersEmployerInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно получить в [списке работодателей](#tag/Rabotodatel/operation/search-employer) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersEmployerInfo**](EmployersEmployerInfo.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerManager

> EmployerManagersEmployerManagerInfo GetEmployerManager(ctx, employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение информации о менеджере

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	managerId := "managerId_example" // string | Идентификатор менеджера
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerManager(context.Background(), employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerManager`: EmployerManagersEmployerManagerInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**managerId** | **string** | Идентификатор менеджера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagersEmployerManagerInfo**](EmployerManagersEmployerManagerInfo.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerManagerLimits

> EmployerManagersEmployerManagerLimits GetEmployerManagerLimits(ctx, employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Дневной лимит просмотра резюме для текущего менеджера



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	managerId := "managerId_example" // string | Идентификатор менеджера, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerManagerLimits(context.Background(), employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerManagerLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerManagerLimits`: EmployerManagersEmployerManagerLimits
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerManagerLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**managerId** | **string** | Идентификатор менеджера, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerManagerLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagersEmployerManagerLimits**](EmployerManagersEmployerManagerLimits.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerManagerTypes

> EmployerManagerTypesResponse GetEmployerManagerTypes(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник типов и прав менеджера

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerManagerTypes(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerManagerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerManagerTypes`: EmployerManagerTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerManagerTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerManagerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagerTypesResponse**](EmployerManagerTypesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerManagerTypes_0

> EmployerManagerTypesResponse GetEmployerManagerTypes_0(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник типов и прав менеджера

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerManagerTypes_0(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerManagerTypes_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerManagerTypes_0`: EmployerManagerTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerManagerTypes_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerManagerTypes_11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagerTypesResponse**](EmployerManagerTypesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerManager_0

> EmployerManagersEmployerManagerInfo GetEmployerManager_0(ctx, employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение информации о менеджере

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	managerId := "managerId_example" // string | Идентификатор менеджера
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerManager_0(context.Background(), employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerManager_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerManager_0`: EmployerManagersEmployerManagerInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerManager_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать в [информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**managerId** | **string** | Идентификатор менеджера | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerManager_12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagersEmployerManagerInfo**](EmployerManagersEmployerManagerInfo.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerManagers

> EmployerManagersResponse GetEmployerManagers(ctx, employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Список менеджеров работодателя

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (optional) (default to 0)
	perPage := float32(8.14) // float32 | Количество элементов (optional) (default to 200)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerManagers(context.Background(), employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerManagers`: EmployerManagersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы | [default to 0]
 **perPage** | **float32** | Количество элементов | [default to 200]
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagersResponse**](EmployerManagersResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerManagers_0

> EmployerManagersResponse GetEmployerManagers_0(ctx, employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Список менеджеров работодателя

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (optional) (default to 0)
	perPage := float32(8.14) // float32 | Количество элементов (optional) (default to 200)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerManagers_0(context.Background(), employerId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerManagers_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerManagers_0`: EmployerManagersResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerManagers_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerManagers_13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы | [default to 0]
 **perPage** | **float32** | Количество элементов | [default to 200]
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerManagersResponse**](EmployerManagersResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployerVacancyAreas

> EmployersEmployerVacancyAreasResponse GetEmployerVacancyAreas(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список регионов, в которых есть активные вакансии



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetEmployerVacancyAreas(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEmployerVacancyAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployerVacancyAreas`: EmployersEmployerVacancyAreasResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEmployerVacancyAreas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployerVacancyAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersEmployerVacancyAreasResponse**](EmployersEmployerVacancyAreasResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFaculties

> []IncludesIdName GetFaculties(ctx, id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список факультетов учебного заведения



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор учебного заведения, который можно узнать из [подсказки](#tag/Podskazki/operation/get-educational-institutions-suggests)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetFaculties(context.Background(), id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFaculties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFaculties`: []IncludesIdName
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFaculties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор учебного заведения, который можно узнать из [подсказки](#tag/Podskazki/operation/get-educational-institutions-suggests) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFacultiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]IncludesIdName**](IncludesIdName.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoriteVacancies

> VacanciesVacanciesFavoritedResponse GetFavoriteVacancies(ctx).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Список отобранных вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию - 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов на странице (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetFavoriteVacancies(context.Background()).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFavoriteVacancies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoriteVacancies`: VacanciesVacanciesFavoritedResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFavoriteVacancies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoriteVacanciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию - 0) | 
 **perPage** | **float32** | Количество элементов на странице | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacanciesFavoritedResponse**](VacanciesVacanciesFavoritedResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldsOfStudySuggestions

> SuggestsFieldsOfStudy GetFieldsOfStudySuggestions(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по специализациям

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска специализаций. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetFieldsOfStudySuggestions(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFieldsOfStudySuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldsOfStudySuggestions`: SuggestsFieldsOfStudy
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFieldsOfStudySuggestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldsOfStudySuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска специализаций. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsFieldsOfStudy**](SuggestsFieldsOfStudy.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHiddenVacancies

> VacanciesDeletedVacancyListResponse GetHiddenVacancies(ctx, employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()

Список удаленных вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	managerId := "managerId_example" // string | Идентификатор менеджера. Передайте, если требуется получить удаленные вакансии другого менеджера.  Если передать несколько параметров `manager_id`, будет использоваться только последний. По умолчанию возвращаются вакансии текущего пользователя  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий в архиве. Справочник с возможными значениями: `employer_hidden_vacancies_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	perPage := int32(56) // int32 | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение `per_page` составляет 1000  (optional)
	page := int32(56) // int32 | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetHiddenVacancies(context.Background(), employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetHiddenVacancies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHiddenVacancies`: VacanciesDeletedVacancyListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetHiddenVacancies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHiddenVacanciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **managerId** | **string** | Идентификатор менеджера. Передайте, если требуется получить удаленные вакансии другого менеджера.  Если передать несколько параметров &#x60;manager_id&#x60;, будет использоваться только последний. По умолчанию возвращаются вакансии текущего пользователя  | 
 **orderBy** | **string** | Сортировка списка вакансий в архиве. Справочник с возможными значениями: &#x60;employer_hidden_vacancies_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **perPage** | **int32** | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение &#x60;per_page&#x60; составляет 1000  | 
 **page** | **int32** | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesDeletedVacancyListResponse**](VacanciesDeletedVacancyListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHiddenVacancies_0

> VacanciesDeletedVacancyListResponse GetHiddenVacancies_0(ctx, employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()

Список удаленных вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	managerId := "managerId_example" // string | Идентификатор менеджера. Передайте, если требуется получить удаленные вакансии другого менеджера.  Если передать несколько параметров `manager_id`, будет использоваться только последний. По умолчанию возвращаются вакансии текущего пользователя  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий в архиве. Справочник с возможными значениями: `employer_hidden_vacancies_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	perPage := int32(56) // int32 | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение `per_page` составляет 1000  (optional)
	page := int32(56) // int32 | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetHiddenVacancies_0(context.Background(), employerId).HHUserAgent(hHUserAgent).ManagerId(managerId).OrderBy(orderBy).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetHiddenVacancies_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHiddenVacancies_0`: VacanciesDeletedVacancyListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetHiddenVacancies_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHiddenVacancies_14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **managerId** | **string** | Идентификатор менеджера. Передайте, если требуется получить удаленные вакансии другого менеджера.  Если передать несколько параметров &#x60;manager_id&#x60;, будет использоваться только последний. По умолчанию возвращаются вакансии текущего пользователя  | 
 **orderBy** | **string** | Сортировка списка вакансий в архиве. Справочник с возможными значениями: &#x60;employer_hidden_vacancies_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **perPage** | **int32** | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение &#x60;per_page&#x60; составляет 1000  | 
 **page** | **int32** | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesDeletedVacancyListResponse**](VacanciesDeletedVacancyListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndustries

> []DictionariesBranchItem GetIndustries(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Отрасли компаний



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetIndustries(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIndustries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndustries`: []DictionariesBranchItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIndustries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndustriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]DictionariesBranchItem**](DictionariesBranchItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanguages

> []DictionariesLangItem GetLanguages(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список всех языков

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetLanguages(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanguages`: []DictionariesLangItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLanguages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]DictionariesLangItem**](DictionariesLangItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocales

> []LocalesLocaleItem GetLocales(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список доступных локалей



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetLocales(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLocales``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocales`: []LocalesLocaleItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLocales`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]LocalesLocaleItem**](LocalesLocaleItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalesForResume

> []LocalesResumeLocaleItem GetLocalesForResume(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список доступных локалей для резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetLocalesForResume(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLocalesForResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalesForResume`: []LocalesResumeLocaleItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLocalesForResume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalesForResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]LocalesResumeLocaleItem**](LocalesResumeLocaleItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailTemplates

> []MailTemplatesMailTemplate GetMailTemplates(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список доступных шаблонов ответов соискателю



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMailTemplates(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMailTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailTemplates`: []MailTemplatesMailTemplate
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMailTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]MailTemplatesMailTemplate**](MailTemplatesMailTemplate.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagerAccounts

> ManagerAccounts GetManagerAccounts(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Рабочие аккаунты менеджера



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetManagerAccounts(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetManagerAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagerAccounts`: ManagerAccounts
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetManagerAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetManagerAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ManagerAccounts**](ManagerAccounts.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagerSettings

> ManagerSettings GetManagerSettings(ctx, employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Предпочтения менеджера



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "1455" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	managerId := "87654321" // string | Идентификатор менеджера. Можно узнать из списка [менеджеров](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetManagerSettings(context.Background(), employerId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetManagerSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagerSettings`: ManagerSettings
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetManagerSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**managerId** | **string** | Идентификатор менеджера. Можно узнать из списка [менеджеров](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ManagerSettings**](ManagerSettings.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetroStations

> []MetroMetroItem GetMetroStations(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список станций метро во всех городах

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMetroStations(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetroStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetroStations`: []MetroMetroItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetroStations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetroStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]MetroMetroItem**](MetroMetroItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetroStationsInCity

> MetroCityMetroItem GetMetroStationsInCity(ctx, cityId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список станций метро в указанном городе

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	cityId := int32(56) // int32 | Идентификатор города
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMetroStationsInCity(context.Background(), cityId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetroStationsInCity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetroStationsInCity`: MetroCityMetroItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetroStationsInCity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cityId** | **int32** | Идентификатор города | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetroStationsInCityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**MetroCityMetroItem**](MetroCityMetroItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMineResumes

> ResumesMineResponse GetMineResumes(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список резюме авторизованного пользователя

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetMineResumes(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMineResumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMineResumes`: ResumesMineResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMineResumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMineResumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesMineResponse**](ResumesMineResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewResumeConditions

> ResumesResumeConditions GetNewResumeConditions(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Условия заполнения полей нового резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetNewResumeConditions(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNewResumeConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewResumeConditions`: ResumesResumeConditions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNewResumeConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNewResumeConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesResumeConditions**](ResumesResumeConditions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayableApiActions

> EmployerServicesEmployerServices GetPayableApiActions(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация по активным услугам API для платных методов



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := int32(56) // int32 | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetPayableApiActions(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPayableApiActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayableApiActions`: EmployerServicesEmployerServices
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPayableApiActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **int32** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayableApiActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerServicesEmployerServices**](EmployerServicesEmployerServices.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPositionsSuggestions

> SuggestsPositions GetPositionsSuggestions(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по должностям резюме

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "водитель" // string | Текст для поиска должности. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetPositionsSuggestions(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPositionsSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPositionsSuggestions`: SuggestsPositions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPositionsSuggestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPositionsSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска должности. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsPositions**](SuggestsPositions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrefNegotiationsOrder

> VacanciesPreferredNegotiationsOrder GetPrefNegotiationsOrder(ctx, id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Просмотр предпочитаемой сортировки откликов

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetPrefNegotiationsOrder(context.Background(), id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPrefNegotiationsOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrefNegotiationsOrder`: VacanciesPreferredNegotiationsOrder
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPrefNegotiationsOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrefNegotiationsOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesPreferredNegotiationsOrder**](VacanciesPreferredNegotiationsOrder.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfessionalRolesDictionary

> ProfessionalRolesCatalog GetProfessionalRolesDictionary(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник профессиональных ролей



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetProfessionalRolesDictionary(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProfessionalRolesDictionary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfessionalRolesDictionary`: ProfessionalRolesCatalog
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProfessionalRolesDictionary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfessionalRolesDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ProfessionalRolesCatalog**](ProfessionalRolesCatalog.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfessionalRolesSuggests

> SuggestsProfessionalRoles GetProfessionalRolesSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по профессиональным ролям

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "водитель" // string | Текст для поиска профессиональной роли. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetProfessionalRolesSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProfessionalRolesSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfessionalRolesSuggests`: SuggestsProfessionalRoles
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProfessionalRolesSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfessionalRolesSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска профессиональной роли. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsProfessionalRoles**](SuggestsProfessionalRoles.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProlongationVacancyInfo

> VacanciesVacancyProlongate GetProlongationVacancyInfo(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Информация о возможности продления вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetProlongationVacancyInfo(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProlongationVacancyInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProlongationVacancyInfo`: VacanciesVacancyProlongate
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProlongationVacancyInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProlongationVacancyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancyProlongate**](VacanciesVacancyProlongate.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredCompaniesSuggests

> SuggestsCompanies GetRegisteredCompaniesSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по зарегистрированным организациям



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска организации. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetRegisteredCompaniesSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRegisteredCompaniesSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegisteredCompaniesSuggests`: SuggestsCompanies
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRegisteredCompaniesSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredCompaniesSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска организации. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsCompanies**](SuggestsCompanies.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResume

> ResumeResumeViewResponse GetResume(ctx, resumeId).HHUserAgent(hHUserAgent).WithNegotiationsHistory(withNegotiationsHistory).WithCreds(withCreds).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()

Просмотр резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	withNegotiationsHistory := true // bool | В случае, если передан данный параметр, в ответе добавится поле negotiations_history.vacancies.  Его формат подробно описан в методе [полной истории откликов/приглашений по резюме](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-resume-negotiations-history) и различается лишь тем,  что в данном случае список будет ограничен тремя вакансиями данного работодателя и последним изменением состояния отклика/приглашения по каждой из этих вакансий  (optional)
	withCreds := true // bool | В случае, если передан данный параметр, в ответе добавится поле creds (optional)
	withJobSearchStatus := true // bool | Параметр для просмотра в резюме статуса поиска кандидата  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResume(context.Background(), resumeId).HHUserAgent(hHUserAgent).WithNegotiationsHistory(withNegotiationsHistory).WithCreds(withCreds).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResume`: ResumeResumeViewResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **withNegotiationsHistory** | **bool** | В случае, если передан данный параметр, в ответе добавится поле negotiations_history.vacancies.  Его формат подробно описан в методе [полной истории откликов/приглашений по резюме](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-resume-negotiations-history) и различается лишь тем,  что в данном случае список будет ограничен тремя вакансиями данного работодателя и последним изменением состояния отклика/приглашения по каждой из этих вакансий  | 
 **withCreds** | **bool** | В случае, если передан данный параметр, в ответе добавится поле creds | 
 **withJobSearchStatus** | **bool** | Параметр для просмотра в резюме статуса поиска кандидата  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumeResumeViewResponse**](ResumeResumeViewResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeAccessTypes

> ResumesAccessTypes GetResumeAccessTypes(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение списка типов видимости резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeAccessTypes(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeAccessTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeAccessTypes`: ResumesAccessTypes
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeAccessTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeAccessTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesAccessTypes**](ResumesAccessTypes.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeConditions

> ResumesResumeConditions GetResumeConditions(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Условия заполнения полей существующего резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeConditions(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeConditions`: ResumesResumeConditions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesResumeConditions**](ResumesResumeConditions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeCreationAvailability

> ResumesCreationAvailability GetResumeCreationAvailability(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Проверка возможности создания резюме

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeCreationAvailability(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeCreationAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeCreationAvailability`: ResumesCreationAvailability
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeCreationAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeCreationAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesCreationAvailability**](ResumesCreationAvailability.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeNegotiationsHistory

> ResumesResumeNegotiationsHistoryResponse GetResumeNegotiationsHistory(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

История откликов/приглашений по резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeNegotiationsHistory(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeNegotiationsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeNegotiationsHistory`: ResumesResumeNegotiationsHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeNegotiationsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeNegotiationsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesResumeNegotiationsHistoryResponse**](ResumesResumeNegotiationsHistoryResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeSearchKeywordsSuggests

> SuggestsSearchKeyword GetResumeSearchKeywordsSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по ключевым словам поиска резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска ключевого слова. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeSearchKeywordsSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeSearchKeywordsSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeSearchKeywordsSuggests`: SuggestsSearchKeyword
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeSearchKeywordsSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeSearchKeywordsSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска ключевого слова. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsSearchKeyword**](SuggestsSearchKeyword.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeStatus

> ResumeStatusReadiness GetResumeStatus(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Статус резюме и готовность к публикации



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeStatus(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeStatus`: ResumeStatusReadiness
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumeStatusReadiness**](ResumeStatusReadiness.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeViewHistory

> ResumesResumeViewHistoryResponse GetResumeViewHistory(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

История просмотра резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeViewHistory(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeViewHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeViewHistory`: ResumesResumeViewHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeViewHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeViewHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesResumeViewHistoryResponse**](ResumesResumeViewHistoryResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeVisibilityEmployersList

> ResumesResumeVisibilityListSearchResponse GetResumeVisibilityEmployersList(ctx, resumeId, listType).Text(text).HHUserAgent(hHUserAgent).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()

Поиск работодателей для добавления в список видимости



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	listType := "listType_example" // string | Тип списка. Допустимые значения — `whitelist` или `blacklist`
	text := "text_example" // string | Строка для поиска. Переданное значение ищется в начале названия работодателя и в начале названия департаментов работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	perPage := int32(56) // int32 | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение per_page составляет 100  (optional)
	page := int32(56) // int32 | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeVisibilityEmployersList(context.Background(), resumeId, listType).Text(text).HHUserAgent(hHUserAgent).PerPage(perPage).Page(page).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeVisibilityEmployersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeVisibilityEmployersList`: ResumesResumeVisibilityListSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeVisibilityEmployersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 
**listType** | **string** | Тип списка. Допустимые значения — &#x60;whitelist&#x60; или &#x60;blacklist&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeVisibilityEmployersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **text** | **string** | Строка для поиска. Переданное значение ищется в начале названия работодателя и в начале названия департаментов работодателя | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **perPage** | **int32** | Количество элементов на странице выдачи. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). Значение по умолчанию и максимальное значение per_page составляет 100  | 
 **page** | **int32** | Порядковый номер страницы в выдаче. Поддерживаются [стандартные параметры пагинации](https://github.com/hhru/api/blob/master/docs/general.md#pagination). По умолчанию нумерация начинается с 0 страницы  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesResumeVisibilityListSearchResponse**](ResumesResumeVisibilityListSearchResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeVisibilityList

> ResumesGetResumeVisibilityListResponse GetResumeVisibilityList(ctx, resumeId, listType).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение списка видимости резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	listType := "listType_example" // string | Тип списка. Допустимые значения — `whitelist` или `blacklist`
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumeVisibilityList(context.Background(), resumeId, listType).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumeVisibilityList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeVisibilityList`: ResumesGetResumeVisibilityListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumeVisibilityList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 
**listType** | **string** | Тип списка. Допустимые значения — &#x60;whitelist&#x60; или &#x60;blacklist&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeVisibilityListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesGetResumeVisibilityListResponse**](ResumesGetResumeVisibilityListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResume_0

> ResumeResumeViewResponse GetResume_0(ctx, resumeId).HHUserAgent(hHUserAgent).WithNegotiationsHistory(withNegotiationsHistory).WithCreds(withCreds).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()

Просмотр резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	withNegotiationsHistory := true // bool | В случае, если передан данный параметр, в ответе добавится поле negotiations_history.vacancies.  Его формат подробно описан в методе [полной истории откликов/приглашений по резюме](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-resume-negotiations-history) и различается лишь тем,  что в данном случае список будет ограничен тремя вакансиями данного работодателя и последним изменением состояния отклика/приглашения по каждой из этих вакансий  (optional)
	withCreds := true // bool | В случае, если передан данный параметр, в ответе добавится поле creds (optional)
	withJobSearchStatus := true // bool | Параметр для просмотра в резюме статуса поиска кандидата  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResume_0(context.Background(), resumeId).HHUserAgent(hHUserAgent).WithNegotiationsHistory(withNegotiationsHistory).WithCreds(withCreds).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResume_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResume_0`: ResumeResumeViewResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResume_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResume_15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **withNegotiationsHistory** | **bool** | В случае, если передан данный параметр, в ответе добавится поле negotiations_history.vacancies.  Его формат подробно описан в методе [полной истории откликов/приглашений по резюме](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-resume-negotiations-history) и различается лишь тем,  что в данном случае список будет ограничен тремя вакансиями данного работодателя и последним изменением состояния отклика/приглашения по каждой из этих вакансий  | 
 **withCreds** | **bool** | В случае, если передан данный параметр, в ответе добавится поле creds | 
 **withJobSearchStatus** | **bool** | Параметр для просмотра в резюме статуса поиска кандидата  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumeResumeViewResponse**](ResumeResumeViewResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResume_1

> ResumeResumeViewResponse GetResume_1(ctx, resumeId).HHUserAgent(hHUserAgent).WithNegotiationsHistory(withNegotiationsHistory).WithCreds(withCreds).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()

Просмотр резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	withNegotiationsHistory := true // bool | В случае, если передан данный параметр, в ответе добавится поле negotiations_history.vacancies.  Его формат подробно описан в методе [полной истории откликов/приглашений по резюме](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-resume-negotiations-history) и различается лишь тем,  что в данном случае список будет ограничен тремя вакансиями данного работодателя и последним изменением состояния отклика/приглашения по каждой из этих вакансий  (optional)
	withCreds := true // bool | В случае, если передан данный параметр, в ответе добавится поле creds (optional)
	withJobSearchStatus := true // bool | Параметр для просмотра в резюме статуса поиска кандидата  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResume_1(context.Background(), resumeId).HHUserAgent(hHUserAgent).WithNegotiationsHistory(withNegotiationsHistory).WithCreds(withCreds).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResume_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResume_1`: ResumeResumeViewResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResume_1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResume_16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **withNegotiationsHistory** | **bool** | В случае, если передан данный параметр, в ответе добавится поле negotiations_history.vacancies.  Его формат подробно описан в методе [полной истории откликов/приглашений по резюме](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-resume-negotiations-history) и различается лишь тем,  что в данном случае список будет ограничен тремя вакансиями данного работодателя и последним изменением состояния отклика/приглашения по каждой из этих вакансий  | 
 **withCreds** | **bool** | В случае, если передан данный параметр, в ответе добавится поле creds | 
 **withJobSearchStatus** | **bool** | Параметр для просмотра в резюме статуса поиска кандидата  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumeResumeViewResponse**](ResumeResumeViewResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumesByStatus

> ResumesByStatusResponse GetResumesByStatus(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Резюме, сгруппированные по возможности отклика на данную вакансию



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetResumesByStatus(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResumesByStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumesByStatus`: ResumesByStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResumesByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResumesByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesByStatusResponse**](ResumesByStatusResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalaryEmployeeLevels

> []IncludesIdNameDesc GetSalaryEmployeeLevels(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Уровни компетенций



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSalaryEmployeeLevels(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSalaryEmployeeLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalaryEmployeeLevels`: []IncludesIdNameDesc
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSalaryEmployeeLevels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalaryEmployeeLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]IncludesIdNameDesc**](IncludesIdNameDesc.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalaryEvaluation

> SalaryStatisticsEvaluationResponse GetSalaryEvaluation(ctx, areaId).HHUserAgent(hHUserAgent).ExcludeArea(excludeArea).EmployeeLevel(employeeLevel).Industry(industry).Speciality(speciality).ExtendSources(extendSources).PositionName(positionName).Locale(locale).Host(host).Execute()

Оценка заработной платы без прогноза



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	areaId := "areaId_example" // string | Код [региона](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), по которому будет построена выборка для формирования отчета 
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	excludeArea := "excludeArea_example" // string | Коды [регионов](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), которые будут исключены из выборки для формирования отчета. Параметр позволяет получить оценку на региональном рынке за исключением определенных городов или областей  (optional)
	employeeLevel := "employeeLevel_example" // string | Справочник [уровни компетенций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-employee-levels), которые будут включены в выборку для формирования отчета  (optional)
	industry := "industry_example" // string | Справочник [Коды отраслей](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-industries), по которым будет построена выборка для формирования отчета  (optional)
	speciality := "speciality_example" // string | Справочник [Коды профобластей и специализаций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-professional-areas), которые будут включены в выборку для формирования отчета  (optional)
	extendSources := true // bool | Использовать ли данные из резюме и вакансий, если по указанным параметрам не нашлось данных в банке зарплат. По умолчанию — `false`  (optional)
	positionName := "positionName_example" // string | Наименование должности. Если не переданы параметры `speciality` или `employee_level`, сервис самостоятельно определит возможные специализации и уровень специалиста по указанной должности и отрасли, и построит отчет по ним  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSalaryEvaluation(context.Background(), areaId).HHUserAgent(hHUserAgent).ExcludeArea(excludeArea).EmployeeLevel(employeeLevel).Industry(industry).Speciality(speciality).ExtendSources(extendSources).PositionName(positionName).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSalaryEvaluation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalaryEvaluation`: SalaryStatisticsEvaluationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSalaryEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**areaId** | **string** | Код [региона](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), по которому будет построена выборка для формирования отчета  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalaryEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **excludeArea** | **string** | Коды [регионов](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), которые будут исключены из выборки для формирования отчета. Параметр позволяет получить оценку на региональном рынке за исключением определенных городов или областей  | 
 **employeeLevel** | **string** | Справочник [уровни компетенций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-employee-levels), которые будут включены в выборку для формирования отчета  | 
 **industry** | **string** | Справочник [Коды отраслей](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-industries), по которым будет построена выборка для формирования отчета  | 
 **speciality** | **string** | Справочник [Коды профобластей и специализаций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-professional-areas), которые будут включены в выборку для формирования отчета  | 
 **extendSources** | **bool** | Использовать ли данные из резюме и вакансий, если по указанным параметрам не нашлось данных в банке зарплат. По умолчанию — &#x60;false&#x60;  | 
 **positionName** | **string** | Наименование должности. Если не переданы параметры &#x60;speciality&#x60; или &#x60;employee_level&#x60;, сервис самостоятельно определит возможные специализации и уровень специалиста по указанной должности и отрасли, и построит отчет по ним  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SalaryStatisticsEvaluationResponse**](SalaryStatisticsEvaluationResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalaryEvaluation_0

> SalaryStatisticsEvaluationResponse GetSalaryEvaluation_0(ctx, areaId).HHUserAgent(hHUserAgent).ExcludeArea(excludeArea).EmployeeLevel(employeeLevel).Industry(industry).Speciality(speciality).ExtendSources(extendSources).PositionName(positionName).Locale(locale).Host(host).Execute()

Оценка заработной платы без прогноза



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	areaId := "areaId_example" // string | Код [региона](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), по которому будет построена выборка для формирования отчета 
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	excludeArea := "excludeArea_example" // string | Коды [регионов](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), которые будут исключены из выборки для формирования отчета. Параметр позволяет получить оценку на региональном рынке за исключением определенных городов или областей  (optional)
	employeeLevel := "employeeLevel_example" // string | Справочник [уровни компетенций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-employee-levels), которые будут включены в выборку для формирования отчета  (optional)
	industry := "industry_example" // string | Справочник [Коды отраслей](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-industries), по которым будет построена выборка для формирования отчета  (optional)
	speciality := "speciality_example" // string | Справочник [Коды профобластей и специализаций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-professional-areas), которые будут включены в выборку для формирования отчета  (optional)
	extendSources := true // bool | Использовать ли данные из резюме и вакансий, если по указанным параметрам не нашлось данных в банке зарплат. По умолчанию — `false`  (optional)
	positionName := "positionName_example" // string | Наименование должности. Если не переданы параметры `speciality` или `employee_level`, сервис самостоятельно определит возможные специализации и уровень специалиста по указанной должности и отрасли, и построит отчет по ним  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSalaryEvaluation_0(context.Background(), areaId).HHUserAgent(hHUserAgent).ExcludeArea(excludeArea).EmployeeLevel(employeeLevel).Industry(industry).Speciality(speciality).ExtendSources(extendSources).PositionName(positionName).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSalaryEvaluation_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalaryEvaluation_0`: SalaryStatisticsEvaluationResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSalaryEvaluation_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**areaId** | **string** | Код [региона](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), по которому будет построена выборка для формирования отчета  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalaryEvaluation_17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **excludeArea** | **string** | Коды [регионов](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-salary-areas), которые будут исключены из выборки для формирования отчета. Параметр позволяет получить оценку на региональном рынке за исключением определенных городов или областей  | 
 **employeeLevel** | **string** | Справочник [уровни компетенций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-employee-levels), которые будут включены в выборку для формирования отчета  | 
 **industry** | **string** | Справочник [Коды отраслей](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-industries), по которым будет построена выборка для формирования отчета  | 
 **speciality** | **string** | Справочник [Коды профобластей и специализаций](#tag/Spravochniki-Banka-dannyh-zarabotnyh-plat/operation/get-salary-professional-areas), которые будут включены в выборку для формирования отчета  | 
 **extendSources** | **bool** | Использовать ли данные из резюме и вакансий, если по указанным параметрам не нашлось данных в банке зарплат. По умолчанию — &#x60;false&#x60;  | 
 **positionName** | **string** | Наименование должности. Если не переданы параметры &#x60;speciality&#x60; или &#x60;employee_level&#x60;, сервис самостоятельно определит возможные специализации и уровень специалиста по указанной должности и отрасли, и построит отчет по ним  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SalaryStatisticsEvaluationResponse**](SalaryStatisticsEvaluationResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalaryIndustries

> []DictionariesBranchItem GetSalaryIndustries(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Отрасли и сферы деятельности



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSalaryIndustries(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSalaryIndustries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalaryIndustries`: []DictionariesBranchItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSalaryIndustries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalaryIndustriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]DictionariesBranchItem**](DictionariesBranchItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalaryProfessionalAreas

> []DictionariesSalaryStatisticsProfessionalAreasResponseInner GetSalaryProfessionalAreas(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Профобласти и специализации



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSalaryProfessionalAreas(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSalaryProfessionalAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalaryProfessionalAreas`: []DictionariesSalaryStatisticsProfessionalAreasResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSalaryProfessionalAreas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalaryProfessionalAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]DictionariesSalaryStatisticsProfessionalAreasResponseInner**](DictionariesSalaryStatisticsProfessionalAreasResponseInner.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalarySalaryAreas

> []DictionariesSalaryStatisticsAreaItem GetSalarySalaryAreas(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Регионы и города



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSalarySalaryAreas(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSalarySalaryAreas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalarySalaryAreas`: []DictionariesSalaryStatisticsAreaItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSalarySalaryAreas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalarySalaryAreasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**[]DictionariesSalaryStatisticsAreaItem**](DictionariesSalaryStatisticsAreaItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedResumeSearch

> SavedSearchesSavedSearchItem GetSavedResumeSearch(ctx, id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение единичного сохраненного поиска резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор сохраненного поиска из [списка](#tag/Sohranennye-poiski-rezyume/operation/get-saved-resume-searches)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSavedResumeSearch(context.Background(), id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSavedResumeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedResumeSearch`: SavedSearchesSavedSearchItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSavedResumeSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор сохраненного поиска из [списка](#tag/Sohranennye-poiski-rezyume/operation/get-saved-resume-searches) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedResumeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SavedSearchesSavedSearchItem**](SavedSearchesSavedSearchItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedResumeSearches

> SavedSearchesSavedSearchResponse GetSavedResumeSearches(ctx).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Список сохраненных поисков резюме

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию - 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию - 5, максимальное значение - 10) (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSavedResumeSearches(context.Background()).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSavedResumeSearches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedResumeSearches`: SavedSearchesSavedSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSavedResumeSearches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedResumeSearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию - 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию - 5, максимальное значение - 10) | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SavedSearchesSavedSearchResponse**](SavedSearchesSavedSearchResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedVacancySearch

> SavedSearchesSavedSearchItem GetSavedVacancySearch(ctx, id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение единичного сохраненного поиска вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор сохраненного поиска
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSavedVacancySearch(context.Background(), id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSavedVacancySearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedVacancySearch`: SavedSearchesSavedSearchItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSavedVacancySearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор сохраненного поиска | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedVacancySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SavedSearchesSavedSearchItem**](SavedSearchesSavedSearchItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedVacancySearches

> SavedSearchesSavedSearchResponse GetSavedVacancySearches(ctx).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Список сохраненных поисков вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию - 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию - 10, максимальное значение - 10) (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSavedVacancySearches(context.Background()).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSavedVacancySearches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedVacancySearches`: SavedSearchesSavedSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSavedVacancySearches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedVacancySearchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию - 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию - 10, максимальное значение - 10) | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SavedSearchesSavedSearchResponse**](SavedSearchesSavedSearchResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkillSetSuggests

> SuggestsSkillSet GetSkillSetSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по ключевым навыкам

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска ключевых навыков. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSkillSetSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSkillSetSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkillSetSuggests`: SuggestsSkillSet
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSkillSetSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillSetSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска ключевых навыков. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsSkillSet**](SuggestsSkillSet.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkills

> DictionariesSkillsResponse GetSkills(ctx).Id(id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник ключевых навыков



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := int32(56) // int32 | Идентификаторы ключевых навыков. Идентификатор конкретного навыка можно узнать по [подсказке](#tag/Podskazki/operation/get-skill-set-suggests). Передать можно не более 50 значений. Например: `?id=2716&id=3019&id=0`. Если был передан идентификатор несуществующего ключевого навыка, для него не вернется никакой информации
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSkills(context.Background()).Id(id).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkills`: DictionariesSkillsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | Идентификаторы ключевых навыков. Идентификатор конкретного навыка можно узнать по [подсказке](#tag/Podskazki/operation/get-skill-set-suggests). Передать можно не более 50 значений. Например: &#x60;?id&#x3D;2716&amp;id&#x3D;3019&amp;id&#x3D;0&#x60;. Если был передан идентификатор несуществующего ключевого навыка, для него не вернется никакой информации | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**DictionariesSkillsResponse**](DictionariesSkillsResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuitableResumes

> ResumesSuitableResumesResponse GetSuitableResumes(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список подходящих для отклика резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetSuitableResumes(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSuitableResumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuitableResumes`: ResumesSuitableResumesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSuitableResumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuitableResumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesSuitableResumesResponse**](ResumesSuitableResumesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestsDictionary

> EmployerDictionariesTestsResponse GetTestsDictionary(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник тестов работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetTestsDictionary(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTestsDictionary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestsDictionary`: EmployerDictionariesTestsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTestsDictionary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestsDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerDictionariesTestsResponse**](EmployerDictionariesTestsResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTestsDictionary_0

> EmployerDictionariesTestsResponse GetTestsDictionary_0(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Справочник тестов работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetTestsDictionary_0(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTestsDictionary_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestsDictionary_0`: EmployerDictionariesTestsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTestsDictionary_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTestsDictionary_18Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployerDictionariesTestsResponse**](EmployerDictionariesTestsResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancies

> VacanciesVacanciesResponse GetVacancies(ctx).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).AcceptTemporary(acceptTemporary).Locale(locale).Host(host).Execute()

Поиск по вакансиям



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (optional) (default to 0)
	perPage := float32(8.14) // float32 | Количество элементов (optional) (default to 10)
	text := "text_example" // string | Переданное значение ищется в полях вакансии, указанных в параметре `search_field`. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) (optional)
	searchField := "searchField_example" // string | Область поиска. Справочник с возможными значениями: `vacancy_search_fields` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  (optional)
	experience := "experience_example" // string | Опыт работы. Необходимо передавать `id` из справочника `experience` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	employment := "employment_example" // string | Тип занятости. Необходимо передавать `id` из справочника `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	schedule := "schedule_example" // string | График работы. Необходимо передавать `id` из справочника `schedule` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	area := "area_example" // string | Регион. Необходимо передавать `id` из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  (optional)
	metro := "metro_example" // string | Ветка или станция метро. Необходимо передавать `id` из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  (optional)
	professionalRole := "professionalRole_example" // string | Профессиональная область. Необходимо передавать `id` из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  (optional)
	industry := "industry_example" // string | Индустрия компании, разместившей вакансию. Необходимо передавать `id` из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  (optional)
	employerId := "employerId_example" // string | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  (optional)
	currency := "currency_example" // string | Код валюты. Справочник с возможными значениями: `currency` (ключ `code`) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром `salary`  (optional)
	salary := float32(8.14) // float32 | Размер заработной платы. Если указано это поле, но не указано `currency`, то для `currency` используется значение RUR. При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании `salary=100&currency=EUR` будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR. По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте `only_with_salary=true`  (optional)
	label := "label_example" // string | Фильтр по меткам вакансий. Необходимо передавать `id` из справочника `vacancy_label` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	onlyWithSalary := true // bool | Показывать вакансии только с указанием зарплаты. По умолчанию `false`  (optional)
	period := float32(8.14) // float32 | Количество дней, в пределах которых производится поиск по вакансиям  (optional)
	dateFrom := "dateFrom_example" // string | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`. Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	dateTo := "dateTo_example" // string | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`. Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	topLat := float32(8.14) // float32 | Верхняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	bottomLat := float32(8.14) // float32 | Нижняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	leftLng := float32(8.14) // float32 | Левая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	rightLng := float32(8.14) // float32 | Правая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий. Справочник с возможными значениями: `vacancy_search_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Если выбрана сортировка по удалённости от гео-точки `distance`, необходимо также задать её координаты: `sort_point_lat`, `sort_point_lng`  (optional)
	sortPointLat := float32(8.14) // float32 | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	sortPointLng := float32(8.14) // float32 | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	clusters := true // bool | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — `false`  (optional)
	describeArguments := true // bool | Возвращать ли описание использованных параметров поиска (массив `arguments`). По умолчанию — `false`  (optional)
	noMagic := true // bool | Если значение `true` — автоматическое преобразование вакансий отключено. По умолчанию – false. При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос `text=москва бухгалтер 100500` будет преобразован в `text=бухгалтер&only_with_salary=true&area=1&salary=100500`  (optional)
	premium := true // bool | Если значение `true` — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  (optional)
	responsesCountEnabled := true // bool | Если значение `true` — дополнительное поле `counters` с количеством откликов для вакансии включено. По-умолчанию — `false`  (optional)
	partTime := "partTime_example" // string | Вакансии для подработки. Возможные значения: * Все элементы из `working_days` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_intervals` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_modes` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы `part` или `project` из `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент `accept_temporary`, показывает вакансии только с временным трудоустройством. Можно указать несколько значений  (optional)
	acceptTemporary := true // bool | Если значение `true` — то поиск происходит только по вакансиям временной работы. По-умолчанию — `false`  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancies(context.Background()).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).AcceptTemporary(acceptTemporary).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancies`: VacanciesVacanciesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVacanciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы | [default to 0]
 **perPage** | **float32** | Количество элементов | [default to 10]
 **text** | **string** | Переданное значение ищется в полях вакансии, указанных в параметре &#x60;search_field&#x60;. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) | 
 **searchField** | **string** | Область поиска. Справочник с возможными значениями: &#x60;vacancy_search_fields&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  | 
 **experience** | **string** | Опыт работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;experience&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **employment** | **string** | Тип занятости. Необходимо передавать &#x60;id&#x60; из справочника &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **schedule** | **string** | График работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;schedule&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **area** | **string** | Регион. Необходимо передавать &#x60;id&#x60; из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  | 
 **metro** | **string** | Ветка или станция метро. Необходимо передавать &#x60;id&#x60; из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  | 
 **professionalRole** | **string** | Профессиональная область. Необходимо передавать &#x60;id&#x60; из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  | 
 **industry** | **string** | Индустрия компании, разместившей вакансию. Необходимо передавать &#x60;id&#x60; из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  | 
 **employerId** | **string** | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  | 
 **currency** | **string** | Код валюты. Справочник с возможными значениями: &#x60;currency&#x60; (ключ &#x60;code&#x60;) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром &#x60;salary&#x60;  | 
 **salary** | **float32** | Размер заработной платы. Если указано это поле, но не указано &#x60;currency&#x60;, то для &#x60;currency&#x60; используется значение RUR. При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании &#x60;salary&#x3D;100&amp;currency&#x3D;EUR&#x60; будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR. По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте &#x60;only_with_salary&#x3D;true&#x60;  | 
 **label** | **string** | Фильтр по меткам вакансий. Необходимо передавать &#x60;id&#x60; из справочника &#x60;vacancy_label&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **onlyWithSalary** | **bool** | Показывать вакансии только с указанием зарплаты. По умолчанию &#x60;false&#x60;  | 
 **period** | **float32** | Количество дней, в пределах которых производится поиск по вакансиям  | 
 **dateFrom** | **string** | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;. Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **dateTo** | **string** | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;. Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **topLat** | **float32** | Верхняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **bottomLat** | **float32** | Нижняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **leftLng** | **float32** | Левая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **rightLng** | **float32** | Правая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **orderBy** | **string** | Сортировка списка вакансий. Справочник с возможными значениями: &#x60;vacancy_search_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Если выбрана сортировка по удалённости от гео-точки &#x60;distance&#x60;, необходимо также задать её координаты: &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;  | 
 **sortPointLat** | **float32** | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **sortPointLng** | **float32** | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **clusters** | **bool** | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — &#x60;false&#x60;  | 
 **describeArguments** | **bool** | Возвращать ли описание использованных параметров поиска (массив &#x60;arguments&#x60;). По умолчанию — &#x60;false&#x60;  | 
 **noMagic** | **bool** | Если значение &#x60;true&#x60; — автоматическое преобразование вакансий отключено. По умолчанию – false. При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос &#x60;text&#x3D;москва бухгалтер 100500&#x60; будет преобразован в &#x60;text&#x3D;бухгалтер&amp;only_with_salary&#x3D;true&amp;area&#x3D;1&amp;salary&#x3D;100500&#x60;  | 
 **premium** | **bool** | Если значение &#x60;true&#x60; — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  | 
 **responsesCountEnabled** | **bool** | Если значение &#x60;true&#x60; — дополнительное поле &#x60;counters&#x60; с количеством откликов для вакансии включено. По-умолчанию — &#x60;false&#x60;  | 
 **partTime** | **string** | Вакансии для подработки. Возможные значения: * Все элементы из &#x60;working_days&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_intervals&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_modes&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы &#x60;part&#x60; или &#x60;project&#x60; из &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент &#x60;accept_temporary&#x60;, показывает вакансии только с временным трудоустройством. Можно указать несколько значений  | 
 **acceptTemporary** | **bool** | Если значение &#x60;true&#x60; — то поиск происходит только по вакансиям временной работы. По-умолчанию — &#x60;false&#x60;  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacanciesResponse**](VacanciesVacanciesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacanciesSimilarToResume

> VacanciesVacanciesResponse GetVacanciesSimilarToResume(ctx, resumeId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).Locale(locale).Host(host).Execute()

Поиск по вакансиям, похожим на резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию - 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию - 10, максимальное значение - 100) (optional)
	text := "text_example" // string | Переданное значение ищется в полях вакансии, указанных в параметре `search_field`. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) (optional)
	searchField := "searchField_example" // string | Область поиска. Справочник с возможными значениями: `vacancy_search_fields` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  (optional)
	experience := "experience_example" // string | Опыт работы. Необходимо передавать `id` из справочника `experience` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	employment := "employment_example" // string | Тип занятости. Необходимо передавать `id` из справочника `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	schedule := "schedule_example" // string | График работы. Необходимо передавать `id` из справочника `schedule` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	area := "area_example" // string | Регион. Необходимо передавать `id` из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  (optional)
	metro := "metro_example" // string | Ветка или станция метро. Необходимо передавать `id` из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  (optional)
	professionalRole := "professionalRole_example" // string | Профессиональная область. Необходимо передавать `id` из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  (optional)
	industry := "industry_example" // string | Индустрия компании, разместившей вакансию. Необходимо передавать `id` из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  (optional)
	employerId := "employerId_example" // string | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  (optional)
	currency := "currency_example" // string | Код валюты. Справочник с возможными значениями: `currency` (ключ `code`) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром `salary`  (optional)
	salary := float32(8.14) // float32 | Размер заработной платы. Если указано это поле, но не указано `currency`, то для `currency` используется значение RUR. При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании `salary=100&currency=EUR` будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR. По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте `only_with_salary=true`  (optional)
	label := "label_example" // string | Фильтр по меткам вакансий. Необходимо передавать `id` из справочника `vacancy_label` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	onlyWithSalary := true // bool | Показывать вакансии только с указанием зарплаты. По умолчанию `false`  (optional)
	period := float32(8.14) // float32 | Количество дней, в пределах которых производится поиск по вакансиям  (optional)
	dateFrom := "dateFrom_example" // string | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`. Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	dateTo := "dateTo_example" // string | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`. Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	topLat := float32(8.14) // float32 | Верхняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	bottomLat := float32(8.14) // float32 | Нижняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	leftLng := float32(8.14) // float32 | Левая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	rightLng := float32(8.14) // float32 | Правая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий. Справочник с возможными значениями: `vacancy_search_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Если выбрана сортировка по удалённости от гео-точки `distance`, необходимо также задать её координаты: `sort_point_lat`, `sort_point_lng`  (optional)
	sortPointLat := float32(8.14) // float32 | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	sortPointLng := float32(8.14) // float32 | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	clusters := true // bool | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — `false`  (optional)
	describeArguments := true // bool | Возвращать ли описание использованных параметров поиска. Успешный ответ будет содержать поле [`arguments`]((#tag/Poisk-vakansij/operation/get-vacancies))). По умолчанию — `false`  (optional)
	noMagic := true // bool | Если значение `true` — автоматическое преобразование вакансий отключено. По умолчанию – false. При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос `text=москва бухгалтер 100500` будет преобразован в `text=бухгалтер&only_with_salary=true&area=1&salary=100500`  (optional)
	premium := true // bool | Если значение `true` — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  (optional)
	responsesCountEnabled := true // bool | Если значение `true` — дополнительное поле `counters` с количеством откликов для вакансии включено. По-умолчанию — `false`  (optional)
	partTime := "partTime_example" // string | Вакансии для подработки. Возможные значения: * Все элементы из `working_days` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_intervals` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_modes` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы `part` или `project` из `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент `accept_temporary`, показывает вакансии только с временным трудоустройством. Можно указать несколько значений  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacanciesSimilarToResume(context.Background(), resumeId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacanciesSimilarToResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacanciesSimilarToResume`: VacanciesVacanciesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacanciesSimilarToResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacanciesSimilarToResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию - 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию - 10, максимальное значение - 100) | 
 **text** | **string** | Переданное значение ищется в полях вакансии, указанных в параметре &#x60;search_field&#x60;. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) | 
 **searchField** | **string** | Область поиска. Справочник с возможными значениями: &#x60;vacancy_search_fields&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  | 
 **experience** | **string** | Опыт работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;experience&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **employment** | **string** | Тип занятости. Необходимо передавать &#x60;id&#x60; из справочника &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **schedule** | **string** | График работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;schedule&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **area** | **string** | Регион. Необходимо передавать &#x60;id&#x60; из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  | 
 **metro** | **string** | Ветка или станция метро. Необходимо передавать &#x60;id&#x60; из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  | 
 **professionalRole** | **string** | Профессиональная область. Необходимо передавать &#x60;id&#x60; из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  | 
 **industry** | **string** | Индустрия компании, разместившей вакансию. Необходимо передавать &#x60;id&#x60; из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  | 
 **employerId** | **string** | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  | 
 **currency** | **string** | Код валюты. Справочник с возможными значениями: &#x60;currency&#x60; (ключ &#x60;code&#x60;) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром &#x60;salary&#x60;  | 
 **salary** | **float32** | Размер заработной платы. Если указано это поле, но не указано &#x60;currency&#x60;, то для &#x60;currency&#x60; используется значение RUR. При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании &#x60;salary&#x3D;100&amp;currency&#x3D;EUR&#x60; будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR. По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте &#x60;only_with_salary&#x3D;true&#x60;  | 
 **label** | **string** | Фильтр по меткам вакансий. Необходимо передавать &#x60;id&#x60; из справочника &#x60;vacancy_label&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **onlyWithSalary** | **bool** | Показывать вакансии только с указанием зарплаты. По умолчанию &#x60;false&#x60;  | 
 **period** | **float32** | Количество дней, в пределах которых производится поиск по вакансиям  | 
 **dateFrom** | **string** | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;. Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **dateTo** | **string** | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;. Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **topLat** | **float32** | Верхняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **bottomLat** | **float32** | Нижняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **leftLng** | **float32** | Левая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **rightLng** | **float32** | Правая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **orderBy** | **string** | Сортировка списка вакансий. Справочник с возможными значениями: &#x60;vacancy_search_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Если выбрана сортировка по удалённости от гео-точки &#x60;distance&#x60;, необходимо также задать её координаты: &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;  | 
 **sortPointLat** | **float32** | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **sortPointLng** | **float32** | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **clusters** | **bool** | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — &#x60;false&#x60;  | 
 **describeArguments** | **bool** | Возвращать ли описание использованных параметров поиска. Успешный ответ будет содержать поле [&#x60;arguments&#x60;]((#tag/Poisk-vakansij/operation/get-vacancies))). По умолчанию — &#x60;false&#x60;  | 
 **noMagic** | **bool** | Если значение &#x60;true&#x60; — автоматическое преобразование вакансий отключено. По умолчанию – false. При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос &#x60;text&#x3D;москва бухгалтер 100500&#x60; будет преобразован в &#x60;text&#x3D;бухгалтер&amp;only_with_salary&#x3D;true&amp;area&#x3D;1&amp;salary&#x3D;100500&#x60;  | 
 **premium** | **bool** | Если значение &#x60;true&#x60; — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  | 
 **responsesCountEnabled** | **bool** | Если значение &#x60;true&#x60; — дополнительное поле &#x60;counters&#x60; с количеством откликов для вакансии включено. По-умолчанию — &#x60;false&#x60;  | 
 **partTime** | **string** | Вакансии для подработки. Возможные значения: * Все элементы из &#x60;working_days&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_intervals&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_modes&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы &#x60;part&#x60; или &#x60;project&#x60; из &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент &#x60;accept_temporary&#x60;, показывает вакансии только с временным трудоустройством. Можно указать несколько значений  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacanciesResponse**](VacanciesVacanciesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacanciesSimilarToVacancy

> VacanciesVacanciesResponse GetVacanciesSimilarToVacancy(ctx, vacancyId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).Locale(locale).Host(host).Execute()

Поиск по вакансиям, похожим на вакансию



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (optional) (default to 0)
	perPage := float32(8.14) // float32 | Количество элементов (optional) (default to 10)
	text := "text_example" // string | Переданное значение ищется в полях вакансии, указанных в параметре `search_field`. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) (optional)
	searchField := "searchField_example" // string | Область поиска. Справочник с возможными значениями: `vacancy_search_fields` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  (optional)
	experience := "experience_example" // string | Опыт работы. Необходимо передавать `id` из справочника `experience` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	employment := "employment_example" // string | Тип занятости. Необходимо передавать `id` из справочника `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	schedule := "schedule_example" // string | График работы. Необходимо передавать `id` из справочника `schedule` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	area := "area_example" // string | Регион. Необходимо передавать `id` из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  (optional)
	metro := "metro_example" // string | Ветка или станция метро. Необходимо передавать `id` из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  (optional)
	professionalRole := "professionalRole_example" // string | Профессиональная область. Необходимо передавать `id` из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  (optional)
	industry := "industry_example" // string | Индустрия компании, разместившей вакансию. Необходимо передавать `id` из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  (optional)
	employerId := "employerId_example" // string | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  (optional)
	currency := "currency_example" // string | Код валюты. Справочник с возможными значениями: `currency` (ключ `code`) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром `salary`  (optional)
	salary := float32(8.14) // float32 | Размер заработной платы. Если указано это поле, но не указано `currency`, то для `currency` используется значение RUR. При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании `salary=100&currency=EUR` будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR. По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте `only_with_salary=true`  (optional)
	label := "label_example" // string | Фильтр по меткам вакансий. Необходимо передавать `id` из справочника `vacancy_label` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	onlyWithSalary := true // bool | Показывать вакансии только с указанием зарплаты. По умолчанию `false`  (optional)
	period := float32(8.14) // float32 | Количество дней, в пределах которых производится поиск по вакансиям  (optional)
	dateFrom := "dateFrom_example" // string | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`. Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	dateTo := "dateTo_example" // string | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром `period`. Значение указывается в формате `ISO 8601 - YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Указанное значение будет округлено до ближайших пяти минут  (optional)
	topLat := float32(8.14) // float32 | Верхняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	bottomLat := float32(8.14) // float32 | Нижняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	leftLng := float32(8.14) // float32 | Левая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	rightLng := float32(8.14) // float32 | Правая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка вакансий. Справочник с возможными значениями: `vacancy_search_order` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Если выбрана сортировка по удалённости от гео-точки `distance`, необходимо также задать её координаты: `sort_point_lat`, `sort_point_lng`  (optional)
	sortPointLat := float32(8.14) // float32 | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	sortPointLng := float32(8.14) // float32 | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если `order_by` установлено в `distance`  (optional)
	clusters := true // bool | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — `false`  (optional)
	describeArguments := true // bool | Возвращать ли описание использованных параметров поиска. Успешный ответ будет содержать поле [`arguments`]((#tag/Poisk-vakansij/operation/get-vacancies))). По умолчанию — `false`  (optional)
	noMagic := true // bool | Если значение `true` — автоматическое преобразование вакансий отключено. По умолчанию – false. При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос `text=москва бухгалтер 100500` будет преобразован в `text=бухгалтер&only_with_salary=true&area=1&salary=100500`  (optional)
	premium := true // bool | Если значение `true` — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  (optional)
	responsesCountEnabled := true // bool | Если значение `true` — дополнительное поле `counters` с количеством откликов для вакансии включено. По-умолчанию — `false`  (optional)
	partTime := "partTime_example" // string | Вакансии для подработки. Возможные значения: * Все элементы из `working_days` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_intervals` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из `working_time_modes` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы `part` или `project` из `employment` в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент `accept_temporary`, показывает вакансии только с временным трудоустройством. Можно указать несколько значений  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacanciesSimilarToVacancy(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Text(text).SearchField(searchField).Experience(experience).Employment(employment).Schedule(schedule).Area(area).Metro(metro).ProfessionalRole(professionalRole).Industry(industry).EmployerId(employerId).Currency(currency).Salary(salary).Label(label).OnlyWithSalary(onlyWithSalary).Period(period).DateFrom(dateFrom).DateTo(dateTo).TopLat(topLat).BottomLat(bottomLat).LeftLng(leftLng).RightLng(rightLng).OrderBy(orderBy).SortPointLat(sortPointLat).SortPointLng(sortPointLng).Clusters(clusters).DescribeArguments(describeArguments).NoMagic(noMagic).Premium(premium).ResponsesCountEnabled(responsesCountEnabled).PartTime(partTime).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacanciesSimilarToVacancy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacanciesSimilarToVacancy`: VacanciesVacanciesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacanciesSimilarToVacancy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacanciesSimilarToVacancyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы | [default to 0]
 **perPage** | **float32** | Количество элементов | [default to 10]
 **text** | **string** | Переданное значение ищется в полях вакансии, указанных в параметре &#x60;search_field&#x60;. Доступен [язык запросов](https://hh.ru/article/1175). Специально для этого поля есть [автодополнение](#tag/Podskazki/operation/get-vacancy-search-keywords) | 
 **searchField** | **string** | Область поиска. Справочник с возможными значениями: &#x60;vacancy_search_fields&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию, используются все поля. Можно указать несколько значений  | 
 **experience** | **string** | Опыт работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;experience&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **employment** | **string** | Тип занятости. Необходимо передавать &#x60;id&#x60; из справочника &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **schedule** | **string** | График работы. Необходимо передавать &#x60;id&#x60; из справочника &#x60;schedule&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **area** | **string** | Регион. Необходимо передавать &#x60;id&#x60; из справочника [/areas](#tag/Obshie-spravochniki/operation/get-areas). Можно указать несколько значений  | 
 **metro** | **string** | Ветка или станция метро. Необходимо передавать &#x60;id&#x60; из справочника [/metro](#tag/Obshie-spravochniki/operation/get-metro-stations). Можно указать несколько значений  | 
 **professionalRole** | **string** | Профессиональная область. Необходимо передавать &#x60;id&#x60; из справочника [/professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary)  | 
 **industry** | **string** | Индустрия компании, разместившей вакансию. Необходимо передавать &#x60;id&#x60; из справочника [/industries](#tag/Obshie-spravochniki/operation/get-industries). Можно указать несколько значений  | 
 **employerId** | **string** | Идентификатор [работодателя](#tag/Rabotodatel). Можно указать несколько значений  | 
 **currency** | **string** | Код валюты. Справочник с возможными значениями: &#x60;currency&#x60; (ключ &#x60;code&#x60;) в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Имеет смысл указывать только совместно с параметром &#x60;salary&#x60;  | 
 **salary** | **float32** | Размер заработной платы. Если указано это поле, но не указано &#x60;currency&#x60;, то для &#x60;currency&#x60; используется значение RUR. При указании значения будут найдены вакансии, в которых вилка зарплаты близка к указанной в запросе. При этом значения пересчитываются по текущим курсам ЦБ РФ. Например, при указании &#x60;salary&#x3D;100&amp;currency&#x3D;EUR&#x60; будут найдены вакансии, где вилка зарплаты указана в рублях и после пересчёта в Евро близка к 100 EUR. По умолчанию будут также найдены вакансии, в которых вилка зарплаты не указана, чтобы такие вакансии отфильтровать, используйте &#x60;only_with_salary&#x3D;true&#x60;  | 
 **label** | **string** | Фильтр по меткам вакансий. Необходимо передавать &#x60;id&#x60; из справочника &#x60;vacancy_label&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **onlyWithSalary** | **bool** | Показывать вакансии только с указанием зарплаты. По умолчанию &#x60;false&#x60;  | 
 **period** | **float32** | Количество дней, в пределах которых производится поиск по вакансиям  | 
 **dateFrom** | **string** | Дата, которая ограничивает снизу диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;. Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **dateTo** | **string** | Дата, которая ограничивает сверху диапазон дат публикации вакансий. Нельзя передавать вместе с параметром &#x60;period&#x60;. Значение указывается в формате &#x60;ISO 8601 - YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Указанное значение будет округлено до ближайших пяти минут  | 
 **topLat** | **float32** | Верхняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **bottomLat** | **float32** | Нижняя граница широты. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **leftLng** | **float32** | Левая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **rightLng** | **float32** | Правая граница долготы. При поиске используется значение указанного в вакансии адреса. Принимаемое значение — градусы в виде десятичной дроби. Необходимо передавать одновременно все четыре параметра гео-координат, иначе вернется ошибка  | 
 **orderBy** | **string** | Сортировка списка вакансий. Справочник с возможными значениями: &#x60;vacancy_search_order&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). Если выбрана сортировка по удалённости от гео-точки &#x60;distance&#x60;, необходимо также задать её координаты: &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;  | 
 **sortPointLat** | **float32** | Значение географической широты точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **sortPointLng** | **float32** | Значение географической долготы точки, по расстоянию от которой будут отсортированы вакансии. Необходимо указывать только, если &#x60;order_by&#x60; установлено в &#x60;distance&#x60;  | 
 **clusters** | **bool** | Возвращать ли [кластеры для данного поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij). По умолчанию — &#x60;false&#x60;  | 
 **describeArguments** | **bool** | Возвращать ли описание использованных параметров поиска. Успешный ответ будет содержать поле [&#x60;arguments&#x60;]((#tag/Poisk-vakansij/operation/get-vacancies))). По умолчанию — &#x60;false&#x60;  | 
 **noMagic** | **bool** | Если значение &#x60;true&#x60; — автоматическое преобразование вакансий отключено. По умолчанию – false. При включённом автоматическом преобразовании, будет предпринята попытка изменить текстовый запрос пользователя на набор параметров. Например, запрос &#x60;text&#x3D;москва бухгалтер 100500&#x60; будет преобразован в &#x60;text&#x3D;бухгалтер&amp;only_with_salary&#x3D;true&amp;area&#x3D;1&amp;salary&#x3D;100500&#x60;  | 
 **premium** | **bool** | Если значение &#x60;true&#x60; — в сортировке вакансий будет учтены премиум вакансии. Такая сортировка используется на сайте. По умолчанию — false  | 
 **responsesCountEnabled** | **bool** | Если значение &#x60;true&#x60; — дополнительное поле &#x60;counters&#x60; с количеством откликов для вакансии включено. По-умолчанию — &#x60;false&#x60;  | 
 **partTime** | **string** | Вакансии для подработки. Возможные значения: * Все элементы из &#x60;working_days&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_intervals&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Все элементы из &#x60;working_time_modes&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элементы &#x60;part&#x60; или &#x60;project&#x60; из &#x60;employment&#x60; в [/dictionaries](#tag/Obshie-spravochniki/operation/get-dictionaries). * Элемент &#x60;accept_temporary&#x60;, показывает вакансии только с временным трудоустройством. Можно указать несколько значений  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacanciesResponse**](VacanciesVacanciesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancy

> VacanciesVacancy GetVacancy(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Просмотр вакансии



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancy(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancy`: VacanciesVacancy
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancy**](VacanciesVacancy.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyBrandedTemplatesList

> EmployersVacancyBrandedTemplatesList GetVacancyBrandedTemplatesList(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список доступных бренд шаблонов вакансий работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "12345" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyBrandedTemplatesList(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyBrandedTemplatesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyBrandedTemplatesList`: EmployersVacancyBrandedTemplatesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyBrandedTemplatesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyBrandedTemplatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersVacancyBrandedTemplatesList**](EmployersVacancyBrandedTemplatesList.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyBrandedTemplatesList_0

> EmployersVacancyBrandedTemplatesList GetVacancyBrandedTemplatesList_0(ctx, employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список доступных бренд шаблонов вакансий работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "12345" // string | Идентификатор работодателя
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyBrandedTemplatesList_0(context.Background(), employerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyBrandedTemplatesList_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyBrandedTemplatesList_0`: EmployersVacancyBrandedTemplatesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyBrandedTemplatesList_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyBrandedTemplatesList_19Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersVacancyBrandedTemplatesList**](EmployersVacancyBrandedTemplatesList.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyConditions

> VacanciesVacancyConditions GetVacancyConditions(ctx).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Условия заполнения полей при добавлении и редактировании вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyConditions(context.Background()).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyConditions`: VacanciesVacancyConditions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyConditions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancyConditions**](VacanciesVacancyConditions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyDraft

> VacancyDraftVacancyDraftFull GetVacancyDraft(ctx, draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Получение черновика вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	draftId := "draftId_example" // string | Идентификатор черновика
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyDraft(context.Background(), draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyDraft`: VacancyDraftVacancyDraftFull
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string** | Идентификатор черновика | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacancyDraftVacancyDraftFull**](VacancyDraftVacancyDraftFull.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyDraftList

> VacancyDraftVacancyDraftItems GetVacancyDraftList(ctx).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Получение списка черновиков вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию - 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию - 20, максимальное значение - 50) (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyDraftList(context.Background()).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyDraftList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyDraftList`: VacancyDraftVacancyDraftItems
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyDraftList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyDraftListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию - 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию - 20, максимальное значение - 50) | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacancyDraftVacancyDraftItems**](VacancyDraftVacancyDraftItems.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyPositionsSuggests

> SuggestsVacancyPositions GetVacancyPositionsSuggests(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по должностям вакансий

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "водитель" // string | Текст для поиска должности в вакансии. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyPositionsSuggests(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyPositionsSuggests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyPositionsSuggests`: SuggestsVacancyPositions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyPositionsSuggests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyPositionsSuggestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска должности в вакансии. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsVacancyPositions**](SuggestsVacancyPositions.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancySearchKeywords

> SuggestsSearchKeyword GetVacancySearchKeywords(ctx).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Подсказки по ключевым словам поиска вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	text := "text_example" // string | Текст для поиска ключевого слова. Искомый текст должен быть длиной два или более символа и не более 3 000 символов
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancySearchKeywords(context.Background()).Text(text).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancySearchKeywords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancySearchKeywords`: SuggestsSearchKeyword
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancySearchKeywords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancySearchKeywordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Текст для поиска ключевого слова. Искомый текст должен быть длиной два или более символа и не более 3 000 символов | 
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**SuggestsSearchKeyword**](SuggestsSearchKeyword.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyStats

> VacanciesVacancyStatsResponse GetVacancyStats(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Статистика по вакансии



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyStats(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyStats`: VacanciesVacancyStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancyStatsResponse**](VacanciesVacancyStatsResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyUpgradeList

> VacanciesVacancyUpgradeListResponse GetVacancyUpgradeList(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Список улучшений для вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyUpgradeList(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyUpgradeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyUpgradeList`: VacanciesVacancyUpgradeListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyUpgradeList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyUpgradeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancyUpgradeListResponse**](VacanciesVacancyUpgradeListResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancyVisitors

> VacanciesVisitorsResponse GetVacancyVisitors(ctx, vacancyId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Посмотревшие вакансию



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию — 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию — 20, максимальное значение — 50) (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancyVisitors(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancyVisitors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancyVisitors`: VacanciesVisitorsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancyVisitors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancyVisitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию — 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию — 20, максимальное значение — 50) | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVisitorsResponse**](VacanciesVisitorsResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVacancy_0

> VacanciesVacancy GetVacancy_0(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Просмотр вакансии



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetVacancy_0(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVacancy_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVacancy_0`: VacanciesVacancy
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVacancy_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVacancy_20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacanciesVacancy**](VacanciesVacancy.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideActiveResponse

> HideActiveResponse(ctx, nid).HHUserAgent(hHUserAgent).WithDeclineMessage(withDeclineMessage).Locale(locale).Host(host).Execute()

Скрыть отклик

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	nid := "nid_example" // string | Идентификатор отклика
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	withDeclineMessage := true // bool | Должно ли быть отправлено работодателю сообщение об отказе, по умолчанию `false`. Возможность отправить сообщение об отказе определяется полем `decline_allowed` получаемым при запросе [списка откликов или одного отклика](https://github.com/hhru/api/blob/e2a0ac4e174a6b56272f78348c05958f5db1b392/docs/negotiations.md#get_negotiation)  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.HideActiveResponse(context.Background(), nid).HHUserAgent(hHUserAgent).WithDeclineMessage(withDeclineMessage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HideActiveResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nid** | **string** | Идентификатор отклика | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideActiveResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **withDeclineMessage** | **bool** | Должно ли быть отправлено работодателю сообщение об отказе, по умолчанию &#x60;false&#x60;. Возможность отправить сообщение об отказе определяется полем &#x60;decline_allowed&#x60; получаемым при запросе [списка откликов или одного отклика](https://github.com/hhru/api/blob/e2a0ac4e174a6b56272f78348c05958f5db1b392/docs/negotiations.md#get_negotiation)  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateToken

> InvalidateToken(ctx).Execute()

Инвалидация токена



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.InvalidateToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvalidateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateTokenRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateToken_0

> InvalidateToken_0(ctx).Execute()

Инвалидация токена



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.InvalidateToken_0(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvalidateToken_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateToken_21Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateToken_1

> InvalidateToken_1(ctx).Execute()

Инвалидация токена



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.InvalidateToken_1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvalidateToken_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateToken_22Request struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadArtifact

> ArtifactsArtifactItem LoadArtifact(ctx).HHUserAgent(hHUserAgent).File(file).Type_(type_).Locale(locale).Host(host).Description(description).Execute()

Загрузка артефакта



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	file := os.NewFile(1234, "some_file") // *os.File | Файл изображения
	type_ := "type__example" // string | Тип артефакта: `photo` или `portfolio`
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")
	description := "description_example" // string | Текстовое описание. Имеет смысл для `portfolio` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.LoadArtifact(context.Background()).HHUserAgent(hHUserAgent).File(file).Type_(type_).Locale(locale).Host(host).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LoadArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoadArtifact`: ArtifactsArtifactItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LoadArtifact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoadArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **file** | ***os.File** | Файл изображения | 
 **type_** | **string** | Тип артефакта: &#x60;photo&#x60; или &#x60;portfolio&#x60; | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]
 **description** | **string** | Текстовое описание. Имеет смысл для &#x60;portfolio&#x60; | 

### Return type

[**ArtifactsArtifactItem**](ArtifactsArtifactItem.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveSavedResumeSearch

> MoveSavedResumeSearch(ctx, savedSearchId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Передача сохраненного поиска резюме другому менеджеру

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	savedSearchId := "savedSearchId_example" // string | Идентификатор из [списка сохраненных поисков](#tag/Sohranennye-poiski-rezyume/operation/get-saved-resume-searches)
	managerId := "managerId_example" // string | Идентификатор менеджера, которому надо передать автопоиск ([список менеджеров компании](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers))
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.MoveSavedResumeSearch(context.Background(), savedSearchId, managerId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MoveSavedResumeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savedSearchId** | **string** | Идентификатор из [списка сохраненных поисков](#tag/Sohranennye-poiski-rezyume/operation/get-saved-resume-searches) | 
**managerId** | **string** | Идентификатор менеджера, которому надо передать автопоиск ([список менеджеров компании](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveSavedResumeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNegotiationsTopicsRead

> PostNegotiationsTopicsRead(ctx).HHUserAgent(hHUserAgent).TopicId(topicId).Locale(locale).Host(host).Execute()

Отметить отклики прочитанными

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	topicId := "topicId_example" // string | Идентификаторы откликов
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.PostNegotiationsTopicsRead(context.Background()).HHUserAgent(hHUserAgent).TopicId(topicId).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostNegotiationsTopicsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNegotiationsTopicsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **topicId** | **string** | Идентификаторы откликов | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishResume

> PublishResume(ctx, resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Публикация резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	resumeId := "resumeId_example" // string | Идентификатор резюме
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.PublishResume(context.Background(), resumeId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PublishResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resumeId** | **string** | Идентификатор резюме | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishVacancy

> IncludesId PublishVacancy(ctx).HHUserAgent(hHUserAgent).VacancyCreate(vacancyCreate).IgnoreDuplicates(ignoreDuplicates).Locale(locale).Host(host).Execute()

Публикация вакансии



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	vacancyCreate := map[string][]openapiclient.VacancyCreate{ ... } // VacancyCreate | 
	ignoreDuplicates := true // bool | Форсирование добавления дубликатов (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PublishVacancy(context.Background()).HHUserAgent(hHUserAgent).VacancyCreate(vacancyCreate).IgnoreDuplicates(ignoreDuplicates).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PublishVacancy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishVacancy`: IncludesId
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PublishVacancy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishVacancyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **vacancyCreate** | [**VacancyCreate**](VacancyCreate.md) |  | 
 **ignoreDuplicates** | **bool** | Форсирование добавления дубликатов | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**IncludesId**](IncludesId.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishVacancyFromDraft

> VacancyDraftVacanciesDraftResponse PublishVacancyFromDraft(ctx, draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Публикация вакансии на основе черновика

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	draftId := "1110031" // string | Идентификатор черновика вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PublishVacancyFromDraft(context.Background(), draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PublishVacancyFromDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishVacancyFromDraft`: VacancyDraftVacanciesDraftResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PublishVacancyFromDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string** | Идентификатор черновика вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishVacancyFromDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacancyDraftVacanciesDraftResponse**](VacancyDraftVacanciesDraftResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMailTemplatesItem

> PutMailTemplatesItem(ctx, employerId, templateId).HHUserAgent(hHUserAgent).MailTemplatesMailTemplateInput(mailTemplatesMailTemplateInput).Locale(locale).Host(host).Execute()

Изменение шаблона ответа соискателю



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info)
	templateId := "templateId_example" // string | Идентификатор шаблона для изменения из [списка доступных шаблонов ответов соискателю](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-mail-templates)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	mailTemplatesMailTemplateInput := *openapiclient.NewMailTemplatesMailTemplateInput("Text_example") // MailTemplatesMailTemplateInput | 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.PutMailTemplatesItem(context.Background(), employerId, templateId).HHUserAgent(hHUserAgent).MailTemplatesMailTemplateInput(mailTemplatesMailTemplateInput).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutMailTemplatesItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя, который можно узнать [в информации о текущем пользователе](#tag/Informaciya-o-menedzhere/operation/get-current-user-info) | 
**templateId** | **string** | Идентификатор шаблона для изменения из [списка доступных шаблонов ответов соискателю](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-mail-templates) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMailTemplatesItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **mailTemplatesMailTemplateInput** | [**MailTemplatesMailTemplateInput**](MailTemplatesMailTemplateInput.md) |  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNegotiationsCollectionToNextState

> PutNegotiationsCollectionToNextState(ctx, collection).HHUserAgent(hHUserAgent).TopicId(topicId).Locale(locale).Host(host).AddressId(addressId).Message(message).SendSms(sendSms).Execute()

Действия по откликам/приглашениям



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	collection := "discard" // string | Идентификатор [коллекции](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#states) топиков, в которую будет перенесено состояние отклика
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	topicId := "topicId_example" // string | Идентификаторы топиков
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")
	addressId := "addressId_example" // string | Идентификатор [адреса](https://api.hh.ru/openapi/redoc#tag/Adresa-rabotodatelya), который будет указан в приглашении (optional)
	message := "message_example" // string | Сообщение, которое будет отправлено соискателю на электронную почту. Используйте [шаблоны](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-mail-templates) для получения текстов (optional)
	sendSms := true // bool | Если установлено `true`, соискателю будет отправлено SMS-уведомление о приглашении. Обратите внимание, что в SMS-сообщении используется стандартный текст, изменить его нельзя (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.PutNegotiationsCollectionToNextState(context.Background(), collection).HHUserAgent(hHUserAgent).TopicId(topicId).Locale(locale).Host(host).AddressId(addressId).Message(message).SendSms(sendSms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutNegotiationsCollectionToNextState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | Идентификатор [коллекции](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#states) топиков, в которую будет перенесено состояние отклика | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNegotiationsCollectionToNextStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **topicId** | **string** | Идентификаторы топиков | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]
 **addressId** | **string** | Идентификатор [адреса](https://api.hh.ru/openapi/redoc#tag/Adresa-rabotodatelya), который будет указан в приглашении | 
 **message** | **string** | Сообщение, которое будет отправлено соискателю на электронную почту. Используйте [шаблоны](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-mail-templates) для получения текстов | 
 **sendSms** | **bool** | Если установлено &#x60;true&#x60;, соискателю будет отправлено SMS-уведомление о приглашении. Обратите внимание, что в SMS-сообщении используется стандартный текст, изменить его нельзя | [default to false]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPrefNegotiationsOrder

> PutPrefNegotiationsOrder(ctx, id).HHUserAgent(hHUserAgent).Order(order).Locale(locale).Host(host).Execute()

Изменение предпочитаемой сортировки откликов

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	order := "order_example" // string | Идентификатор типа сортировки. Возможные значения представлены в поле `order_types` [коллекции откликов и приглашений](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#collections_response). Не все типы сортировки из списка доступны для изменения
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.PutPrefNegotiationsOrder(context.Background(), id).HHUserAgent(hHUserAgent).Order(order).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutPrefNegotiationsOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPrefNegotiationsOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **order** | **string** | Идентификатор типа сортировки. Возможные значения представлены в поле &#x60;order_types&#x60; [коллекции откликов и приглашений](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#collections_response). Не все типы сортировки из списка доступны для изменения | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreVacancyFromHidden

> RestoreVacancyFromHidden(ctx, employerId, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Восстановление вакансии из удаленных



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	employerId := "employerId_example" // string | Идентификатор работодателя
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.RestoreVacancyFromHidden(context.Background(), employerId, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestoreVacancyFromHidden``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employerId** | **string** | Идентификатор работодателя | 
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreVacancyFromHiddenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEmployer

> EmployersEmployersList SearchEmployer(ctx).HHUserAgent(hHUserAgent).Text(text).Area(area).Type_(type_).OnlyWithVacancies(onlyWithVacancies).SortBy(sortBy).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()

Поиск работодателя



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	text := "text_example" // string | Текст для поиска. Переданное значение ищется в названии и описании работодателя (optional)
	area := "area_example" // string | Идентификатор региона работодателя, множественный параметр. Идентификаторы регионов можно узнать в [справочнике регионов](#tag/Obshie-spravochniki/operation/get-areas) (optional)
	type_ := "type__example" // string | Тип работодателя, множественный параметр. Разрешенные значения перечислены в [справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле `employer_type` (optional)
	onlyWithVacancies := true // bool | Возвращать только работодателей у которых есть в данный момент открытые вакансии (`true`) или же всех (`false`). По умолчанию — `false` (optional)
	sortBy := "sortBy_example" // string | Сортировка по имени (`by_name`) или по количеству открытых вакансий (`by_vacancies_open`). По умолчанию — `by_name` (optional)
	page := float32(8.14) // float32 | Номер страницы с работодателями (считается от `0`, по умолчанию — `0`) (optional)
	perPage := float32(8.14) // float32 | Количество элементов на страницу (по умолчанию — 20, максимум — 100 ) (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SearchEmployer(context.Background()).HHUserAgent(hHUserAgent).Text(text).Area(area).Type_(type_).OnlyWithVacancies(onlyWithVacancies).SortBy(sortBy).Page(page).PerPage(perPage).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchEmployer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEmployer`: EmployersEmployersList
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchEmployer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEmployerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **text** | **string** | Текст для поиска. Переданное значение ищется в названии и описании работодателя | 
 **area** | **string** | Идентификатор региона работодателя, множественный параметр. Идентификаторы регионов можно узнать в [справочнике регионов](#tag/Obshie-spravochniki/operation/get-areas) | 
 **type_** | **string** | Тип работодателя, множественный параметр. Разрешенные значения перечислены в [справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле &#x60;employer_type&#x60; | 
 **onlyWithVacancies** | **bool** | Возвращать только работодателей у которых есть в данный момент открытые вакансии (&#x60;true&#x60;) или же всех (&#x60;false&#x60;). По умолчанию — &#x60;false&#x60; | 
 **sortBy** | **string** | Сортировка по имени (&#x60;by_name&#x60;) или по количеству открытых вакансий (&#x60;by_vacancies_open&#x60;). По умолчанию — &#x60;by_name&#x60; | 
 **page** | **float32** | Номер страницы с работодателями (считается от &#x60;0&#x60;, по умолчанию — &#x60;0&#x60;) | 
 **perPage** | **float32** | Количество элементов на страницу (по умолчанию — 20, максимум — 100 ) | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**EmployersEmployersList**](EmployersEmployersList.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchForResumes

> ResumesSearchForResumesResponse SearchForResumes(ctx).HHUserAgent(hHUserAgent).Text(text).TextLogic(textLogic).TextField(textField).TextPeriod(textPeriod).AgeFrom(ageFrom).AgeTo(ageTo).Area(area).Relocation(relocation).Period(period).DateFrom(dateFrom).DateTo(dateTo).EducationLevel(educationLevel).Employment(employment).Experience(experience).Skill(skill).Gender(gender).Label(label).Language(language).LanguageLevel(languageLevel).Metro(metro).Currency(currency).SalaryFrom(salaryFrom).SalaryTo(salaryTo).Schedule(schedule).OrderBy(orderBy).Citizenship(citizenship).WorkTicket(workTicket).EducationalInstitution(educationalInstitution).SearchInResponses(searchInResponses).ByTextPrefix(byTextPrefix).DriverLicenseTypes(driverLicenseTypes).VacancyId(vacancyId).Page(page).PerPage(perPage).ProfessionalRole(professionalRole).Folder(folder).IncludeAllFolders(includeAllFolders).JobSearchStatus(jobSearchStatus).Resume(resume).FilterExpIndustry(filterExpIndustry).FilterExpPeriod(filterExpPeriod).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()

Поиск резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	text := "text_example" // string | Поисковая фраза. Метод найдет резюме, в которых встречаются все слова заданной фразы.  Особенности:  * Можно указать несколько значений. Каждое дополнительное значение уточняет поиск. * В качестве поисковой фразы можно использовать [язык поисковых запросов](http://hh.ru/article.xml?articleId=1175). * Специально для этого поля предусмотрено [автодополнение по подсказкам](#tag/Podskazki/operation/get-resume-search-keywords-suggests). * Для тонкой настройки по фразе можно использовать параметры `text.logic`, `text.field`, `text.period`. При использовании дополнительных `text.*` полей, необходимо указывать весь набор (триаду) параметров  (optional)
	textLogic := "textLogic_example" // string | Описывает, как производится поиск. Возможные значения перечислены в поле `resume_search_logic` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) (optional)
	textField := "textField_example" // string | Описывает, где должны встречаться слова из поисковой фразы `text`. Можно указать несколько значений через запятую, например `?text.field=education,keywords`. Возможные значения перечислены в поле `resume_search_fields` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) (optional)
	textPeriod := "textPeriod_example" // string | Период опыта работы.  Параметр имеет смысл только при `text.field` равным одному из значений: `experience`, `experience_company`, `experience_position`, `experience_description`, но указывать его необходимо всегда при указании других `text.*`. Если параметр не имеет смысла, то его значение можно оставить пустым  (optional)
	ageFrom := "ageFrom_example" // string | Нижняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение `only_with_age` в параметре `label`  (optional)
	ageTo := "ageTo_example" // string | Верхняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение `only_with_age` в параметре `label`  (optional)
	area := "area_example" // string | Регион. Возможные значения указаны в [справочнике регионов](https://github.com/hhru/api/blob/master/docs/areas.md). Можно указать несколько значений.  По умолчанию выбираются резюме, в которых соискатели живут в указанных регионах или готовы в них переехать. Поменять это поведение поиска можно, указав параметр `relocation`  (optional)
	relocation := "relocation_example" // string | Готовность к переезду. Возможные значения указаны в поле `resume_search_relocation` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Необходимо указывать вместе с параметром `area`  (optional)
	period := float32(8.14) // float32 | Поиск ведется по резюме, опубликованным за указанный период в днях. Если период не указан, поиск ведется без ограничений по дате публикации  (optional)
	dateFrom := "dateFrom_example" // string | Дата, от которой нужно начать поиск. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — `YYYY-MM-DD` или с точностью до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Нельзя передавать вместе с параметром `period`  (optional)
	dateTo := "dateTo_example" // string | Дата, до которой нужно искать. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — `YYYY-MM-DD` или с точность до секунды `YYYY-MM-DDThh:mm:ss±hhmm`. Можно передавать только в паре с параметром `date_from`. Нельзя передавать вместе с параметром `period`  (optional)
	educationLevel := "educationLevel_example" // string | Уровень образования. Возможные значения перечислены в поле `education_level` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Если параметр не указан, поиск ведется без ограничений на уровень образования  (optional)
	employment := "employment_example" // string | Тип занятости. Возможные значения перечислены в поле `employment` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	experience := "experience_example" // string | Опыт работы. Возможные значения перечислены в поле `experience` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	skill := "skill_example" // string | Ключевые навыки. Указывается один или несколько идентификаторов ключевых навыков. Значения можно получить из поля `id` в [подсказке по ключевым навыкам](#tag/Podskazki/operation/get-skill-set-suggests)  (optional)
	gender := "gender_example" // string | Пол соискателя. Возможные значения перечислены в поле `gender` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries).  По умолчанию вне зависимости от значения параметра будут найдены резюме, в которых пол не указан, исключить из поисковой выдачи такие резюме можно с помощью параметра `label=only_with_gender`  (optional)
	label := "label_example" // string | Дополнительный фильтр. Возможные значения перечислены в поле `resume_search_label` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	language := "language_example" // string | Знание языка. Можно указать несколько значений.  Возможные значения перечислены в [справочнике языков](#tag/Obshie-spravochniki/operation/get-languages)  (optional)
	languageLevel := "languageLevel_example" // string | Уровень знания языка. Можно указать несколько значений.  Возможные значения перечислены в поле `language_level` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	metro := "metro_example" // string | Линия, либо станция метро. Возможные значения перечислены в [справочнике метро](#tag/Obshie-spravochniki/operation/get-metro-stations)  (optional)
	currency := "currency_example" // string | Код валюты. Возможные значения перечислены в поле `currency.code` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	salaryFrom := float32(8.14) // float32 | Нижняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр `label=only_with_salary`  (optional)
	salaryTo := float32(8.14) // float32 | Верхняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр `label=only_with_salary`  (optional)
	schedule := "schedule_example" // string | График работы. Возможные значения перечислены в поле `schedule` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	orderBy := "orderBy_example" // string | Сортировка списка резюме. Возможные значения перечислены в поле `resume_search_order` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	citizenship := "citizenship_example" // string | Страна гражданства соискателя. Возможные значения перечислены в [справочнике стран](https://github.com/hhru/api/blob/master/docs/areas.md#countries). Можно указать несколько значений  (optional)
	workTicket := "workTicket_example" // string | Страна, в которой у соискателя есть разрешение на работу. Возможные значения перечислены в [справочнике стран](https://github.com/hhru/api/blob/master/docs/areas.md#countries). Можно указать несколько значений  (optional)
	educationalInstitution := "educationalInstitution_example" // string | Учебные заведения соискателя. В качестве параметров используются [подсказки по названиям университетов](#tag/Podskazki/operation/get-educational-institutions-suggests). Можно указать несколько значений  (optional)
	searchInResponses := true // bool | Если `true`, то поиск осуществляется только по резюме, которыми соискатели откликались на вакансии компании текущего пользователя, если `false` — поиск осуществляется по всем резюме. По умолчанию — `false`  (optional)
	byTextPrefix := true // bool | Если `true`, включается поиск по префиксу. Для каждого параметра `text` будут находиться не только полные совпадения слов, но еще и слова, начинающиеся с `text`. По умолчанию — `false`  (optional)
	driverLicenseTypes := "driverLicenseTypes_example" // string | Категории водительских прав соискателя. Возможные значения перечислены в поле `driver_license_types` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  (optional)
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии для поиска похожих резюме. Необходимо передавать идентификатор активной вакансии работодателя или вакансии работодателя в архиве  (optional)
	page := float32(8.14) // float32 | Номер страницы (считается от 0, по умолчанию — 0) (optional)
	perPage := float32(8.14) // float32 | Количество элементов (по умолчанию — 20, максимальное значение — 100) (optional)
	professionalRole := "professionalRole_example" // string | Профессиональная роль. Элемент справочника [профессиональных ролей](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary). Можно указать несколько значений  (optional)
	folder := "folder_example" // string | Один или несколько идентификаторов папок с отобранными резюме. Если данный параметр передан, поиск будет ограничен содержимым указанных папок. Можно передавать идентификаторы нескольких папок, например: `folder=111&folder=222&folder=333`  (optional)
	includeAllFolders := true // bool | Признак, указывающий, нужно ли вести поиск по всем папкам с отобранными резюме.  Если у менеджера есть доступ к избранным папкам, то поиск проходит по умолчанию в избранных папках. Если передать параметр `false`, то поиск не будет ограничен папками. Если в одном запросе будут переданы параметры `folder` и `include_all_folders`, вернется ошибка `400 Bad Request`  (optional)
	jobSearchStatus := "jobSearchStatus_example" // string | Статус поиска работы.  Возможные значения перечислены в поле `job_search_status` в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  (optional)
	resume := "resume_example" // string | Идентификатор резюме для поиска похожих резюме  (optional)
	filterExpIndustry := "filterExpIndustry_example" // string | Обрабатывается совместно с параметром `filter_exp_period`. Идентификатор отрасли, в которой у соискателя должен присутствовать опыт работы. Возможные значения перечислены в [справочнике отраслей](#tag/Obshie-spravochniki/operation/get-industries) (поле id). Можно указать несколько значений  (optional)
	filterExpPeriod := "filterExpPeriod_example" // string | Период, за который у соискателя должен присутствовать опыт работы в отрасли, указанной в параметре `filter_exp_industry`.  Возможные значения:   * `all_time` - за все время * `last_year` - за последний год * `last_three_years` - за последние 3 года * `last_six_years` - за последние 6 лет.  По умолчанию равен `all_time`  (optional)
	withJobSearchStatus := true // bool | Параметр для просмотра в резюме статуса поиска кандидата  (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SearchForResumes(context.Background()).HHUserAgent(hHUserAgent).Text(text).TextLogic(textLogic).TextField(textField).TextPeriod(textPeriod).AgeFrom(ageFrom).AgeTo(ageTo).Area(area).Relocation(relocation).Period(period).DateFrom(dateFrom).DateTo(dateTo).EducationLevel(educationLevel).Employment(employment).Experience(experience).Skill(skill).Gender(gender).Label(label).Language(language).LanguageLevel(languageLevel).Metro(metro).Currency(currency).SalaryFrom(salaryFrom).SalaryTo(salaryTo).Schedule(schedule).OrderBy(orderBy).Citizenship(citizenship).WorkTicket(workTicket).EducationalInstitution(educationalInstitution).SearchInResponses(searchInResponses).ByTextPrefix(byTextPrefix).DriverLicenseTypes(driverLicenseTypes).VacancyId(vacancyId).Page(page).PerPage(perPage).ProfessionalRole(professionalRole).Folder(folder).IncludeAllFolders(includeAllFolders).JobSearchStatus(jobSearchStatus).Resume(resume).FilterExpIndustry(filterExpIndustry).FilterExpPeriod(filterExpPeriod).WithJobSearchStatus(withJobSearchStatus).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchForResumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchForResumes`: ResumesSearchForResumesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchForResumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchForResumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **text** | **string** | Поисковая фраза. Метод найдет резюме, в которых встречаются все слова заданной фразы.  Особенности:  * Можно указать несколько значений. Каждое дополнительное значение уточняет поиск. * В качестве поисковой фразы можно использовать [язык поисковых запросов](http://hh.ru/article.xml?articleId&#x3D;1175). * Специально для этого поля предусмотрено [автодополнение по подсказкам](#tag/Podskazki/operation/get-resume-search-keywords-suggests). * Для тонкой настройки по фразе можно использовать параметры &#x60;text.logic&#x60;, &#x60;text.field&#x60;, &#x60;text.period&#x60;. При использовании дополнительных &#x60;text.*&#x60; полей, необходимо указывать весь набор (триаду) параметров  | 
 **textLogic** | **string** | Описывает, как производится поиск. Возможные значения перечислены в поле &#x60;resume_search_logic&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
 **textField** | **string** | Описывает, где должны встречаться слова из поисковой фразы &#x60;text&#x60;. Можно указать несколько значений через запятую, например &#x60;?text.field&#x3D;education,keywords&#x60;. Возможные значения перечислены в поле &#x60;resume_search_fields&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
 **textPeriod** | **string** | Период опыта работы.  Параметр имеет смысл только при &#x60;text.field&#x60; равным одному из значений: &#x60;experience&#x60;, &#x60;experience_company&#x60;, &#x60;experience_position&#x60;, &#x60;experience_description&#x60;, но указывать его необходимо всегда при указании других &#x60;text.*&#x60;. Если параметр не имеет смысла, то его значение можно оставить пустым  | 
 **ageFrom** | **string** | Нижняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение &#x60;only_with_age&#x60; в параметре &#x60;label&#x60;  | 
 **ageTo** | **string** | Верхняя граница возраста соискателя в годах.  По умолчанию в выдачу добавляются также резюме с неуказанным возрастом. Для выдачи резюме только с указанным возрастом передайте значение &#x60;only_with_age&#x60; в параметре &#x60;label&#x60;  | 
 **area** | **string** | Регион. Возможные значения указаны в [справочнике регионов](https://github.com/hhru/api/blob/master/docs/areas.md). Можно указать несколько значений.  По умолчанию выбираются резюме, в которых соискатели живут в указанных регионах или готовы в них переехать. Поменять это поведение поиска можно, указав параметр &#x60;relocation&#x60;  | 
 **relocation** | **string** | Готовность к переезду. Возможные значения указаны в поле &#x60;resume_search_relocation&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Необходимо указывать вместе с параметром &#x60;area&#x60;  | 
 **period** | **float32** | Поиск ведется по резюме, опубликованным за указанный период в днях. Если период не указан, поиск ведется без ограничений по дате публикации  | 
 **dateFrom** | **string** | Дата, от которой нужно начать поиск. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — &#x60;YYYY-MM-DD&#x60; или с точностью до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Нельзя передавать вместе с параметром &#x60;period&#x60;  | 
 **dateTo** | **string** | Дата, до которой нужно искать. Значение указывается в формате [ISO 8601](https://github.com/hhru/api/blob/master/docs/general.md#date-format) — &#x60;YYYY-MM-DD&#x60; или с точность до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;. Можно передавать только в паре с параметром &#x60;date_from&#x60;. Нельзя передавать вместе с параметром &#x60;period&#x60;  | 
 **educationLevel** | **string** | Уровень образования. Возможные значения перечислены в поле &#x60;education_level&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Если параметр не указан, поиск ведется без ограничений на уровень образования  | 
 **employment** | **string** | Тип занятости. Возможные значения перечислены в поле &#x60;employment&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **experience** | **string** | Опыт работы. Возможные значения перечислены в поле &#x60;experience&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **skill** | **string** | Ключевые навыки. Указывается один или несколько идентификаторов ключевых навыков. Значения можно получить из поля &#x60;id&#x60; в [подсказке по ключевым навыкам](#tag/Podskazki/operation/get-skill-set-suggests)  | 
 **gender** | **string** | Пол соискателя. Возможные значения перечислены в поле &#x60;gender&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries).  По умолчанию вне зависимости от значения параметра будут найдены резюме, в которых пол не указан, исключить из поисковой выдачи такие резюме можно с помощью параметра &#x60;label&#x3D;only_with_gender&#x60;  | 
 **label** | **string** | Дополнительный фильтр. Возможные значения перечислены в поле &#x60;resume_search_label&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **language** | **string** | Знание языка. Можно указать несколько значений.  Возможные значения перечислены в [справочнике языков](#tag/Obshie-spravochniki/operation/get-languages)  | 
 **languageLevel** | **string** | Уровень знания языка. Можно указать несколько значений.  Возможные значения перечислены в поле &#x60;language_level&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **metro** | **string** | Линия, либо станция метро. Возможные значения перечислены в [справочнике метро](#tag/Obshie-spravochniki/operation/get-metro-stations)  | 
 **currency** | **string** | Код валюты. Возможные значения перечислены в поле &#x60;currency.code&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **salaryFrom** | **float32** | Нижняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр &#x60;label&#x3D;only_with_salary&#x60;  | 
 **salaryTo** | **float32** | Верхняя граница желаемой заработной платы (ЗП).  По умолчанию в выдачу добавляются также резюме с неуказанной ЗП. Для выдачи резюме только с указанной ЗП передайте параметр &#x60;label&#x3D;only_with_salary&#x60;  | 
 **schedule** | **string** | График работы. Возможные значения перечислены в поле &#x60;schedule&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **orderBy** | **string** | Сортировка списка резюме. Возможные значения перечислены в поле &#x60;resume_search_order&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **citizenship** | **string** | Страна гражданства соискателя. Возможные значения перечислены в [справочнике стран](https://github.com/hhru/api/blob/master/docs/areas.md#countries). Можно указать несколько значений  | 
 **workTicket** | **string** | Страна, в которой у соискателя есть разрешение на работу. Возможные значения перечислены в [справочнике стран](https://github.com/hhru/api/blob/master/docs/areas.md#countries). Можно указать несколько значений  | 
 **educationalInstitution** | **string** | Учебные заведения соискателя. В качестве параметров используются [подсказки по названиям университетов](#tag/Podskazki/operation/get-educational-institutions-suggests). Можно указать несколько значений  | 
 **searchInResponses** | **bool** | Если &#x60;true&#x60;, то поиск осуществляется только по резюме, которыми соискатели откликались на вакансии компании текущего пользователя, если &#x60;false&#x60; — поиск осуществляется по всем резюме. По умолчанию — &#x60;false&#x60;  | 
 **byTextPrefix** | **bool** | Если &#x60;true&#x60;, включается поиск по префиксу. Для каждого параметра &#x60;text&#x60; будут находиться не только полные совпадения слов, но еще и слова, начинающиеся с &#x60;text&#x60;. По умолчанию — &#x60;false&#x60;  | 
 **driverLicenseTypes** | **string** | Категории водительских прав соискателя. Возможные значения перечислены в поле &#x60;driver_license_types&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
 **vacancyId** | **string** | Идентификатор вакансии для поиска похожих резюме. Необходимо передавать идентификатор активной вакансии работодателя или вакансии работодателя в архиве  | 
 **page** | **float32** | Номер страницы (считается от 0, по умолчанию — 0) | 
 **perPage** | **float32** | Количество элементов (по умолчанию — 20, максимальное значение — 100) | 
 **professionalRole** | **string** | Профессиональная роль. Элемент справочника [профессиональных ролей](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary). Можно указать несколько значений  | 
 **folder** | **string** | Один или несколько идентификаторов папок с отобранными резюме. Если данный параметр передан, поиск будет ограничен содержимым указанных папок. Можно передавать идентификаторы нескольких папок, например: &#x60;folder&#x3D;111&amp;folder&#x3D;222&amp;folder&#x3D;333&#x60;  | 
 **includeAllFolders** | **bool** | Признак, указывающий, нужно ли вести поиск по всем папкам с отобранными резюме.  Если у менеджера есть доступ к избранным папкам, то поиск проходит по умолчанию в избранных папках. Если передать параметр &#x60;false&#x60;, то поиск не будет ограничен папками. Если в одном запросе будут переданы параметры &#x60;folder&#x60; и &#x60;include_all_folders&#x60;, вернется ошибка &#x60;400 Bad Request&#x60;  | 
 **jobSearchStatus** | **string** | Статус поиска работы.  Возможные значения перечислены в поле &#x60;job_search_status&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries). Можно указать несколько значений  | 
 **resume** | **string** | Идентификатор резюме для поиска похожих резюме  | 
 **filterExpIndustry** | **string** | Обрабатывается совместно с параметром &#x60;filter_exp_period&#x60;. Идентификатор отрасли, в которой у соискателя должен присутствовать опыт работы. Возможные значения перечислены в [справочнике отраслей](#tag/Obshie-spravochniki/operation/get-industries) (поле id). Можно указать несколько значений  | 
 **filterExpPeriod** | **string** | Период, за который у соискателя должен присутствовать опыт работы в отрасли, указанной в параметре &#x60;filter_exp_industry&#x60;.  Возможные значения:   * &#x60;all_time&#x60; - за все время * &#x60;last_year&#x60; - за последний год * &#x60;last_three_years&#x60; - за последние 3 года * &#x60;last_six_years&#x60; - за последние 6 лет.  По умолчанию равен &#x60;all_time&#x60;  | 
 **withJobSearchStatus** | **bool** | Параметр для просмотра в резюме статуса поиска кандидата  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumesSearchForResumesResponse**](ResumesSearchForResumesResponse.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchForVacancyDraftDuplicates

> VacancyDuplicates SearchForVacancyDraftDuplicates(ctx, draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Проверка наличия дубликатов вакансии

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	draftId := "12358" // string | Идентификатор черновика вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SearchForVacancyDraftDuplicates(context.Background(), draftId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchForVacancyDraftDuplicates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchForVacancyDraftDuplicates`: VacancyDuplicates
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchForVacancyDraftDuplicates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string** | Идентификатор черновика вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchForVacancyDraftDuplicatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**VacancyDuplicates**](VacancyDuplicates.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCodeForVerifyPhoneInResume

> ResumePhoneGenerateCodeGenerateCode SendCodeForVerifyPhoneInResume(ctx).HHUserAgent(hHUserAgent).Phone(phone).Locale(locale).Host(host).Execute()

Отправить код подтверждения для телефона резюме

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	phone := "phone_example" // string | Телефон на который надо отправить подтверждающий код
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SendCodeForVerifyPhoneInResume(context.Background()).HHUserAgent(hHUserAgent).Phone(phone).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SendCodeForVerifyPhoneInResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendCodeForVerifyPhoneInResume`: ResumePhoneGenerateCodeGenerateCode
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SendCodeForVerifyPhoneInResume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendCodeForVerifyPhoneInResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **phone** | **string** | Телефон на который надо отправить подтверждающий код | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**ResumePhoneGenerateCodeGenerateCode**](ResumePhoneGenerateCodeGenerateCode.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendNegotiationMessage

> NegotiationsMessageSent SendNegotiationMessage(ctx, nid).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Отправка нового сообщения



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	nid := "nid_example" // string | Идентификатор отклика
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.SendNegotiationMessage(context.Background(), nid).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SendNegotiationMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendNegotiationMessage`: NegotiationsMessageSent
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SendNegotiationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nid** | **string** | Идентификатор отклика | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendNegotiationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

[**NegotiationsMessageSent**](NegotiationsMessageSent.md)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicantComment

> UpdateApplicantComment(ctx, applicantId, commentId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Text(text).AccessType(accessType).Execute()

Обновление комментария



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	applicantId := "applicantId_example" // string | Идентификатор соискателя, который можно узнать из поля `owner` [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field)
	commentId := "commentId_example" // string | Идентификатор комментария, который можно узнать в [списке комментариев](#tag/Kommentarii-k-soiskatelyu/operation/get-applicant-comments-list)
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")
	text := "text_example" // string | Текст комментария (optional)
	accessType := "accessType_example" // string | Тип доступа. Доступные значения перечислены [в справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле `applicant_comment_access_type` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.UpdateApplicantComment(context.Background(), applicantId, commentId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Text(text).AccessType(accessType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplicantComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** | Идентификатор соискателя, который можно узнать из поля &#x60;owner&#x60; [в резюме](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#owner-field) | 
**commentId** | **string** | Идентификатор комментария, который можно узнать в [списке комментариев](#tag/Kommentarii-k-soiskatelyu/operation/get-applicant-comments-list) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicantCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]
 **text** | **string** | Текст комментария | 
 **accessType** | **string** | Тип доступа. Доступные значения перечислены [в справочнике](#tag/Obshie-spravochniki/operation/get-dictionaries) в поле &#x60;applicant_comment_access_type&#x60; | 

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSavedResumeSearch

> UpdateSavedResumeSearch(ctx, id).HHUserAgent(hHUserAgent).Name(name).Subscription(subscription).Locale(locale).Host(host).Execute()

Обновление сохраненного поиска резюме



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор сохраненного поиска
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	name := "name_example" // string | Новое имя сохраненного поиска (optional)
	subscription := true // bool | Статус подписки (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.UpdateSavedResumeSearch(context.Background(), id).HHUserAgent(hHUserAgent).Name(name).Subscription(subscription).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSavedResumeSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор сохраненного поиска | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSavedResumeSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **name** | **string** | Новое имя сохраненного поиска | 
 **subscription** | **bool** | Статус подписки | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSavedVacancySearch

> UpdateSavedVacancySearch(ctx, id).HHUserAgent(hHUserAgent).Name(name).Subscription(subscription).Locale(locale).Host(host).Execute()

Обновление сохраненного поиска вакансий



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	id := "id_example" // string | Идентификатор сохраненного поиска
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	name := "name_example" // string | Новое имя сохраненного поиска (optional)
	subscription := true // bool | Статус подписки (optional)
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.UpdateSavedVacancySearch(context.Background(), id).HHUserAgent(hHUserAgent).Name(name).Subscription(subscription).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSavedVacancySearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Идентификатор сохраненного поиска | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSavedVacancySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **name** | **string** | Новое имя сохраненного поиска | 
 **subscription** | **bool** | Статус подписки | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VacancyProlongation

> VacancyProlongation(ctx, vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()

Продление вакансии



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/zaboal/hh-go/hh"
)

func main() {
	vacancyId := "vacancyId_example" // string | Идентификатор вакансии
	hHUserAgent := "MyApp/1.0 (my-app-feedback@example.com)" // string | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC)) 
	locale := "EN" // string | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  (optional) (default to "RU")
	host := "hh.ru" // string | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  (optional) (default to "hh.ru")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.VacancyProlongation(context.Background(), vacancyId).HHUserAgent(hHUserAgent).Locale(locale).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VacancyProlongation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vacancyId** | **string** | Идентификатор вакансии | 

### Other Parameters

Other parameters are passed through a pointer to a apiVacancyProlongationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hHUserAgent** | **string** | Название приложения и контактная почта разработчика (см. [Информация о клиенте](https://github.com/hhru/api/blob/master/docs/general.md#%D1%82%D1%80%D0%B5%D0%B1%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D1%8F-%D0%BA-%D0%B7%D0%B0%D0%BF%D1%80%D0%BE%D1%81%D0%B0%D0%BC))  | 
 **locale** | **string** | Идентификатор локали (см. [Локализация](#tag/Obshie-spravochniki/operation/get-locales))  | [default to &quot;RU&quot;]
 **host** | **string** | Доменное имя сайта (см. [Выбор сайта](https://github.com/hhru/api/blob/master/docs/hosts.md))  | [default to &quot;hh.ru&quot;]

### Return type

 (empty response body)

### Authorization

[OauthToken](../README.md#OauthToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

