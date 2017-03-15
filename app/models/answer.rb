class Answer < ApplicationRecord
  has_many :question_answers
  has_many :questions, :through => :question_answers

  has_many :user_answers
  has_many :users, :through => :user_answers
end
