# Domain models
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
