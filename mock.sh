mockgen -source="handlers/http/responses/success_analysis_response.go" -destination="mocks/success_analysis_response_mock.go" -package=mocks
mockgen -source="analyzers/analyzer.go" -destination="mocks/analyzer_mock.go" -package=mocks
mockgen -source="analyzers/schema/analyzer_info.go" -destination="mocks/analyzer_info_mock.go" -package=mocks
mockgen -source="channels/url_executor.go" -destination="mocks/url_executor_mock.go" -package=mocks
mockgen -source="handlers/http/controllers/body_extractor.go" -destination="mocks/body_extractor_mock.go" -package=mocks
mockgen -source="channels/url_executor_provider.go" -destination="mocks/url_executor_provider_mock.go" -package=mocks
