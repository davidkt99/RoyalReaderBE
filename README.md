# RoyalReaderBE
 The Backend of the RoyalRoad reader made for personal use.

# Queries

Add num of chapters to getBooks endpoint/db call

# Mutations

downloadChapter     return  success
downloadBook        return  success


# Cron Jobs

Add Email user to cronJob


# Database

Add users_books Table
Add users       Table

# Add Users

add user_books  table   user_books_id -> Fkey books_id
add user        table   user_id -> Fkey user_books_id

? user_books      has last read chapter to indicate new chapters    then mutation to update this status




# later
favoriteBook        return  success     fav the book in database

maybe add different users

Add url to book query



# future improvements
more specific sql queries to save on amount of queries
    getBooks for app should have num of chapters thus will need queryAllBook and queryNumOfChapters


password -> https://stackoverflow.com/questions/18656528/how-do-i-encrypt-passwords-with-postgresql
https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72