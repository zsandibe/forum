CREATE TABLE IF NOT EXISTS users(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE,
			Email TEXT NOT NULL,
			Password TEXT
			-- UNIQUE(Email)
);
CREATE TABLE IF NOT EXISTS sessions(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			UserID INTEGER NOT NULL UNIQUE,
			Token VARCHAR(32) NOT NULL,
			ExpDate DATATIME NOT NULL,
			FOREIGN KEY(UserID) REFERENCES USERS(ID)
				ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS posts(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			AuthorID INTEGER NOT NULL,
			Title TEXT NOT NULL,
			Tag TEXT,
			Body TEXT NOT NULL,
			Author TEXT,
			FOREIGN KEY(AuthorID) REFERENCES USERS(ID)
				ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS comments(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			AuthorID INTEGER NOT NULL,
			PostID INTEGER NOT NULL,
			Content TEXT NOT NULL,
			Author TEXT,
			FOREIGN KEY(AuthorID) REFERENCES USERS(ID)
				ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS reactions(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			UserID INTEGER NOT NULL,
			PostID INTEGER,
			CommentID INTEGER,
			VOTE BLOB NOT NULL,
			FOREIGN KEY(UserID) REFERENCES USERS(ID)
				ON DELETE CASCADE,
			FOREIGN KEY(PostID) REFERENCES POSTS(ID)
				ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS tags(
			ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			PostID INTEGER NOT NULL,
			Tag TEXT,
			FOREIGN KEY(PostID) REFERENCES POSTS(ID)
				ON DELETE CASCADE
);
