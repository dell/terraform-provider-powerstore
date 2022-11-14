
# powerstore_snapshot_rule (Resource)

Use this resource type to manage snapshot rules that are used in protection policies.

## Schema

### Snapshot Rule creation Attributes

1. **name**
    * string
    * snapshot rule name
2. **interval**
    * string
    * Interval between snapshots taken by a snapshot rule
    * allowed values :
        * Five_Minutes, Fifteen_Minutes, Thirty_Minutes, One_Hour, Two_Hours, Three_Hours, Four_Hours, Six_Hours, Eight_Hours, Twelve_Hours, One_Day
3. **time_of_day**
    * string
    * Time of the day to take a daily snapshot, with format "hh:mm" using a 24 hour clock.
    * Either the `interval` parameter or the `time_of_day` parameter will be set, but not both
4. **timezone**
    * string
    * Time zone identifier for applying the time zone to the time_of_day for a snapshot rule, including any DST effects if applicable.
    * Applies only when a time_of_day is specified in the snapshot rule.
    * Defaults to UTC if not specified
    * allowed values:
        * Etc__GMT_plus_12, US__Samoa, Etc__GMT_plus_11, America__Atka, US__Hawaii, Etc__GMT_plus_10, Pacific__Marquesas, US__Alaska, Pacific__Gambier, Etc__GMT_plus_9, PST8PDT, Pacific__Pitcairn, US__Pacific, Etc__GMT_plus_8, Mexico__BajaSur, America__Boise, America__Phoenix, MST7MDT, Etc__GMT_plus_7, CST6CDT, America__Chicago, Canada__Saskatchewan, America__Bahia_Banderas, Etc__GMT_plus_6, Chile__EasterIsland, America__Bogota, America__New_York, EST5EDT, America__Havana, Etc__GMT_plus_5, America__Caracas, America__Cuiaba, America__Santo_Domingo, Canada__Atlantic, America__Asuncion, Etc__GMT_plus_4, Canada__Newfoundland, Chile__Continental, Brazil__East, America__Godthab, America__Miquelon, America__Buenos_Aires, Etc__GMT_plus_3, America__Noronha, Etc__GMT_plus_2, America__Scoresbysund, Atlantic__Cape_Verde, Etc__GMT_plus_1, UTC, Europe__London, Africa__Casablanca, Atlantic__Reykjavik, Antarctica__Troll, Europe__Paris, Europe__Sarajevo, Europe__Belgrade, Europe__Rome, Africa__Tunis, Etc__GMT_minus_1, Asia__Gaza, Europe__Bucharest, Europe__Helsinki, Asia__Beirut, Africa__Harare, Asia__Damascus, Asia__Amman, Europe__Tiraspol, Asia__Jerusalem, Etc__GMT_minus_2, Asia__Baghdad, Africa__Asmera, Etc__GMT_minus_3, Asia__Tehran, Asia__Baku, Etc__GMT_minus_4, Asia__Kabul, Asia__Karachi, Etc__GMT_minus_5, Asia__Kolkata, Asia__Katmandu, Asia__Almaty, Etc__GMT_minus_6, Asia__Rangoon, Asia__Hovd, Asia__Bangkok, Etc__GMT_minus_7, Asia__Hong_Kong, Asia__Brunei, Asia__Singapore, Etc__GMT_minus_8, Asia__Pyongyang, Australia__Eucla, Asia__Seoul, Etc__GMT_minus_9, Australia__Darwin, Australia__Adelaide, Australia__Sydney, Australia__Brisbane, Asia__Magadan, Etc__GMT_minus_10, Australia__Lord_Howe, Etc__GMT_minus_11, Asia__Kamchatka, Pacific__Fiji, Antarctica__South_Pole, Etc__GMT_minus_12, Pacific__Chatham, Pacific__Tongatapu, Pacific__Apia, Etc__GMT_minus_13, Pacific__Kiritimati, Etc__GMT_minus_14
5. **days_of_week**
    * [string]
    * Days of the week when the snapshot rule should be applied.
    * Days are determined based on the UTC time zone, unless the time_of_day and timezone properties are set.
    * allowed values:
        * Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday |
6. **DesiredRetention**
    * int32
    * Desired snapshot retention period in hours.
    * The system will retain snapshots for this time period
7. **NASAccessType**
    * NAS filesystem snapshot access method.
    * The setting is ignored for volume, virtual_volume, and volume_group snapshots.
    * allowed values :
        1. Snapshot
            * the files within the snapshot may be access directly from the production file system in the .snapshot subdirectory of each directory.
        2. Protocol
            * the entire file system snapshot may be shared and mounted on a client like any other file system, except that it is readonly


### Snapshot Rule other Attributes
8. **ID**
    * string
    * Unique identifier of the snapshot rule.
9. **IsReplica**
    * bool
    * Indicates whether this is a replica of a snapshot rule on a remote system that is the source of a replication session replicating a storage resource to the local system.
10. **IsReadOnly**
    * bool
    * Indicates whether this snapshot rule can be modified.
11. **ManagedBy**
    * string
    * Entity that owns and manages this instance. The possible values are:
        * User - This instance is managed by the end user.
        * Metro - This instance is managed by the peer system where the policy was assigned, in a Metro Cluster configuration.
        * Replication - This destination instance is managed by the source system in a Replication configuration.
        * VMware_vSphere - This instance is managed by the system through VMware vSphere/vCenter.
    * default: User
12. **ManagedByID**
    * string
    * Unique identifier of the managing entity based on the value of the managed_by property, as shown below:
        * User - Empty
        * Metro - Unique identifier of the remote system where the policy was assigned.
        * Replication - Unique identifier of the source remote system.
        * VMware_vSphere - Unique identifier of the owning VMware vSphere/vCenter