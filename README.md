# RoyalReaderBE
 The Backend of the RoyalRoad reader made for personal use.

# Queries

getBooks


# Mutations

downloadChapter     return  success
downloadBook        return  success


# Cron Jobs

Add Email user to cronJob


# Database

remove read from chapters

# Add Users

add user_books  table   user_books_id -> Fkey books_id
add user        table   user_id -> Fkey user_books_id

? user_books      has last read chapter to indicate new chapters    then mutation to update this status




# later
favoriteBook        return  success     fav the book in database

maybe add different users





# future improvements
more specific sql queries to save on amount of queries
    getBooks for app should have num of chapters thus will need queryAllBook and queryNumOfChapters