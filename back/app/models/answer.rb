class Answer < ApplicationRecord
  belongs_to :question

  has_many :user_answers
  has_many :users, :through => :user_answers

  validates :code, :lang, presence: true
end
