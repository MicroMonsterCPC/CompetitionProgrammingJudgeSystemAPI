class Question < ApplicationRecord
  has_many :answers
  validates :title, :body, presence: true
  validates :title, length: { miximum: 30, minimum: 1 }, uniqueness: true
end
