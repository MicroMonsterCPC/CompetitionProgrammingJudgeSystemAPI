class Answer < ApplicationRecord
  belongs_to :user
  belongs_to :question

  validates :code, :lang, presence: true
end
