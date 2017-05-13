class Question < ApplicationRecord
  has_many :question_answers
  has_many :answers, :through => :question_answers
  validates :title, :body, presence: true
  validates :title, length: { miximum: 30, minimum: 1 }, uniqeness: true
end
