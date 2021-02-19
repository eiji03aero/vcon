# vcon
- matching service for conversational class with vtuber

## Description
### Glossary
- end user
  - user who wants to have lesson
- teacher
  - user who gives lesson
- customer support
  - administration user
- points
  - what is consumed when have a class

### Features
#### end user
##### signup
- with email, google account
##### have a lesson
- end user browse through teachers
  - search with:
    - language
    - avg rating
- end user opens teacher profile
- end user offers lesson on available time
- teacher gets notified, confirm lesson
  - payment thing should be done here
    - subtract points
  - in what way teacher gets notified?
- end user and teacher gets notified before lesson time
- blah blah (no detail implementation for lesson yet)
- end user and teacher both can leave rating
##### purchase points
- end user opens point purchase page
- selects the amount to purchase
- submit, done
##### ask for help from customer support
- just make it an email for now

#### teacher
##### registers available schedule
- teacher opens schedule registration page
- register available schedule
  - calendar to select specific date
  - day-of-week repetitive schedule
##### teacher confirms lesson offer
- teacher receives an email with link to offer confirmation page
- opens confirmation page and confirm lesson
- both end user and teacher receives notification on email

### Domain models
- User
  - id: string @primary
  - userProfile: UserProfile
  - offerEntries: []OfferEntry
  - scheduleEntries: []ScheduleEntry
- UserProfile
  - id: string @primary
  - userId string
  - firstName: string
  - lastName: string
  - description: string
  - profession: Profession
- ScheduleEntry
  - id: string @primary
- OfferEntry
  - id: string @primary
- PurchaseHistory
  - id: string @primary
  - userId: string

### Enums
- Profession
  - student: 1
  - teacher: 2

### Project components
#### web application server
- tools to use
  - golang
    - because I want to learn it
- hexagonal arch
- production environment
  - heroku, docker

#### client application
- tools to use
  - react, nextjs
  - typescript
  - apollo
- hexagonal arch
- production environment
  - static build with netlify

#### database
- postgres

## Todo
- [x] list out what this thing does
- [x] define models
- [x] define what components it has
- [x] start implementing

## Notes
- should be i18n compatible
  - should have ja, en

## References
- heroku
  - https://devcenter.heroku.com/articles/container-registry-and-runtime
