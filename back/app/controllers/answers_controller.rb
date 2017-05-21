class AnswerController < ApplicationController
  def create
    @answer = Answer.new(answer_params)
    if @answer.save
      data = {
        "id": question_id.to_s,
        "code": @answer.code,
        "lang": @answer.lang
      }
      client = Faraday.new(:url => "http://api:1323")
      res = client.post do |req|
        req.url '/answer-data'
        req.headers['Content-Type'] = 'application/json'
        req.body = data.to_json
      end
      return JSON.parse(res.body)["Result"]
    else
      puts "Error"
    end
  end
  private
  def answer_params
    params.requre(:answer).permit(:code, :lang)
  end
end
